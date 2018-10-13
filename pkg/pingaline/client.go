// Copyright © 2018 The Pingaling Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pingaline

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// DefaultBaseURL is where pingaling expects API calls
const DefaultBaseURL = "http://localhost:4000/api"

// Client interacts with the pingaling API
type Client struct {
	HTTPClient http.Client
	BaseURL    string
}

// ErrNotExpectedJSON is returned when the API response isn't expected JSON
type ErrNotExpectedJSON struct {
	OriginalBody string
	Err          error
}

// ErrBadStatusCode is returned when the API returns a non 200 error code
type ErrBadStatusCode struct {
	OriginalBody string
	Code         int
}

func (e *ErrBadStatusCode) Error() string {
	return fmt.Sprintf("Invalid status code: %d", e.Code)
}

func (e *ErrNotExpectedJSON) Error() string {
	return fmt.Sprintf("Unexpected JSON: %s from %s", e.Err.Error(), e.OriginalBody)
}

// CreateSession is a required for further API use.
func (c *Client) CreateSession() (*Session, error) {
	var v createSessionResp
	return &Session{
		parent:    c,
		SessionID: v.SessionID,
	}, nil
}

func (c *Client) urlBase(endpoint string) string {
	base := c.BaseURL
	if c.BaseURL == "" {
		base = DefaultBaseURL
	}
	return fmt.Sprintf("%s/%s", base, endpoint)
}

func (c *Client) doReqURL(ctx context.Context, url string, jsonInto interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := withCancel(ctx, &c.HTTPClient, req)
	if err != nil {
		return err
	}

	defer func() {
		var maxCopySize int64
		maxCopySize = 2 << 10
		io.CopyN(ioutil.Discard, resp.Body, maxCopySize) // drain the response body
		resp.Body.Close()                                //close body
	}()
	var b bytes.Buffer
	if _, err := io.Copy(&b, resp.Body); err != nil {
		return err
	}
	debug := b.String()
	if resp.StatusCode != http.StatusOK {
		return &ErrBadStatusCode{
			OriginalBody: debug,
			Code:         resp.StatusCode,
		}
	}
	if err := json.NewDecoder(&b).Decode(jsonInto); err != nil {
		return &ErrNotExpectedJSON{
			OriginalBody: debug,
			Err:          err,
		}
	}
	return nil
}

func withCancel(ctx context.Context, client *http.Client, req *http.Request) (resp *http.Response, err error) {
	req.Cancel = ctx.Done()
	return client.Do(req)
}
