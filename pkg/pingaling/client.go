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

package pingaling

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

func (c *Client) doReqURL(ctx context.Context, method string, url string,
	headers map[string]string, body io.Reader) (statusCode int, b bytes.Buffer) {

	// Prepare request
	req, err := http.NewRequest(method, url, body)
	CheckError(err)
	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	// make request
	resp, err := withCancel(ctx, &c.HTTPClient, req)
	CheckError(err)

	// drain and close the response body before return
	defer func() {
		var maxCopySize int64
		maxCopySize = 2 << 10
		io.CopyN(ioutil.Discard, resp.Body, maxCopySize)
		resp.Body.Close()
	}()

	// HTTP response Status Code
	statusCode = resp.StatusCode

	// Make a copy of the response body
	_, err = io.Copy(&b, resp.Body)
	CheckError(err)

	return

}

func withCancel(ctx context.Context, client *http.Client, req *http.Request) (resp *http.Response, err error) {
	req.Cancel = ctx.Done()
	return client.Do(req)
}

type RequestServiceGET interface {
	Get(context.Context, string, interface{}) error
}

type RequestServiceDELETE interface {
	Delete(context.Context, string, interface{}) error
}

type RequestServicePOST interface {
	Post(context.Context, string, io.Reader) (bytes.Buffer, error)
}

// Get request
func (c *Client) Get(ctx context.Context, url string, ts interface{}) error {
	statusCode, b := c.doReqURL(ctx, http.MethodGet, url, nil, nil)
	if statusCode != http.StatusOK {
		return &ErrBadStatusCode{
			OriginalBody: b.String(),
			Code:         statusCode,
		}
	}
	// Decode response into target struct
	if err := json.NewDecoder(&b).Decode(ts); err != nil {
		return &ErrNotExpectedJSON{
			OriginalBody: b.String(),
			Err:          err,
		}
	}
	return nil
}

// Delete request
func (c *Client) Delete(ctx context.Context, url string, ts interface{}) error {
	statusCode, b := c.doReqURL(ctx, http.MethodDelete, url, nil, nil)
	if statusCode != http.StatusOK {
		return &ErrBadStatusCode{
			OriginalBody: b.String(),
			Code:         statusCode,
		}
	}
	// Decode response into target struct
	if err := json.NewDecoder(&b).Decode(ts); err != nil {
		return &ErrNotExpectedJSON{
			OriginalBody: b.String(),
			Err:          err,
		}
	}
	return nil
}

// Post request
func (c *Client) Post(ctx context.Context, url string, body io.Reader) (bytes.Buffer, error) {

	headers := make(map[string]string)
	//headers["Content-Type"] = "multipart/mixed; boundary=plug_conn_test"
	headers["Content-Type"] = "application/json"

	statusCode, b := c.doReqURL(ctx, http.MethodPost, url, headers, body)
	if statusCode != http.StatusCreated {
		return bytes.Buffer{}, &ErrBadStatusCode{
			OriginalBody: b.String(),
			Code:         statusCode,
		}
	}
	return b, nil
}
