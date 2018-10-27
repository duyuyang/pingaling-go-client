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
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

const (
	Incidents            = "incidents"
	NotificationChannels = "notification_channels"
	NotificationPolicies = "notification_policies"
	Endpoints            = "endpoints"
	Message              = "Message: "
	HealthSummary        = "health/summary"
	Manifest             = "manifest"
)

// Session establish connection to API
type Session struct {
	parent      *Client
	SessionID   string
	HTTPService HTTPService
}

func (s *Session) url(endpoint string) string {
	u := s.parent.urlBase(endpoint)
	return u
}

// GetHealthStatus return Health check data
func (s *Session) GetHealthStatus() (*HealthData, error) {
	var r HealthData

	b, err := s.HTTPService.Get(s.url(HealthSummary))
	if err != nil {
		return nil, errors.Wrap(err, "GetHealthStatus Get request failed")
	}
	err = JSONDecoder(b, &r)
	if err != nil {
		return nil, errors.Wrap(err, "GetHealthStatus failed to decode JSON")
	}

	return &r, nil

}

// GetEndpoints return specific endpoint data
func (s *Session) GetEndpoints(epName string) (*EndpointData, error) {

	var r EndpointData

	b, err := s.HTTPService.Get(s.url(Endpoints + "/" + epName))
	if err != nil {
		return nil, errors.Wrap(err, "GetEndpoint Get request failed")
	}
	err = JSONDecoder(b, &r)
	if err != nil {
		return nil, errors.Wrap(err, "GetEndpoints failed to decode JSON")
	}
	return &r, nil

}

// GetIncidents return specific endpoint data
func (s *Session) GetIncidents() (*IncidentData, error) {

	var r IncidentData

	b, err := s.HTTPService.Get(s.url(Incidents))
	if err != nil {
		return nil, errors.Wrap(err, "GetIncidents Get request failed")
	}
	err = JSONDecoder(b, &r)
	if err != nil {
		return nil, errors.Wrap(err, "GetIncidents failed to decode JSON")
	}
	return &r, nil

}

// GetNotificationChannels return specific endpoint data
func (s *Session) GetNotificationChannels() (*NotificationChannelData, error) {

	var r NotificationChannelData

	b, err := s.HTTPService.Get(s.url(NotificationChannels))
	if err != nil {
		return nil, errors.Wrap(err, "GetNotificationChannels Get request failed")
	}
	err = JSONDecoder(b, &r)
	if err != nil {
		return nil, errors.Wrap(err, "GetNotificationChannels failed to decode JSON")
	}
	return &r, nil

}

// GetNotificationPolicies return specific endpoint data
func (s *Session) GetNotificationPolicies() (*NotificationPolicyData, error) {

	var r NotificationPolicyData

	b, err := s.HTTPService.Get(s.url(NotificationPolicies))
	if err != nil {
		return nil, errors.Wrap(err, "GetNotificationPolicies Get request failed")
	}
	err = JSONDecoder(b, &r)
	if err != nil {
		return nil, errors.Wrap(err, "GetNotificationPolicies failed to decode JSON")
	}
	return &r, nil

}

// DeleteEndpoints delete specific endpoint
func (s *Session) DeleteEndpoints(name []string) {

	pather := func(i interface{}) interface{} {
		return Endpoints + "/" + i.(string)
	}
	s.deleteIter(pather, name)

}

// DeleteNotificationChannels delete specific notification channels
func (s *Session) DeleteNotificationChannels(name []string) {

	pather := func(i interface{}) interface{} {
		return NotificationChannels + "/" + i.(string)
	}
	s.deleteIter(pather, name)

}

// DeleteNotificationPolicies delete specific notification policies
func (s *Session) DeleteNotificationPolicies(name []string) {

	pather := func(i interface{}) interface{} {
		return NotificationPolicies + "/" + i.(string)
	}
	s.deleteIter(pather, name)

}

func (s *Session) deleteIter(pather func(i interface{}) interface{}, name []string) {

	for i := range Map(pather, StrIter(name)) {
		if m, err := s.deleter(i); err == nil {
			fmt.Println(Message, m)
		} else {
			fmt.Printf("Failed to delete %v", i.(string))
		}
	}

}

func (s *Session) deleter(p interface{}) (interface{}, error) {
	var r DeleteMsg

	b, err := s.HTTPService.Delete(s.url(p.(string)))
	if err != nil {
		return nil, errors.Wrap(err, "deleter Delete request failed")
	}
	err = JSONDecoder(b, &r)
	if err != nil {
		return nil, errors.Wrap(err, "deleter failed to decode JSON")
	}
	return r.Message, nil

}

// ApplyManifest post manifest to API to create resource
func (s *Session) ApplyManifest(doc TypeMeta) (bytes.Buffer, error) {

	manifest := ManifestReq{
		Manifest: doc,
	}
	buff, err := json.Marshal(&manifest)
	if err != nil {
		return bytes.Buffer{}, &ErrNotExpectedJSON{
			OriginalBody: string(buff),
			Err:          err,
		}
	}
	r, err := s.HTTPService.Post(s.url(Manifest), bytes.NewBuffer(buff))
	if err != nil {
		return bytes.Buffer{}, errors.Wrap(err, "ApplyManifest Post request failed")
	}
	return r, nil

}
