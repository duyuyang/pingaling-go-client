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
	"time"
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := s.HTTPService.Get(ctx, s.url(HealthSummary), &r); err != nil {
		return nil, err
	}
	return &r, nil
}

// GetEndpoints return specific endpoint data
func (s *Session) GetEndpoints(epName string) (*EndpointData, error) {

	var r EndpointData
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := s.HTTPService.Get(ctx, s.url(Endpoints+"/"+epName), &r); err != nil {
		return nil, err
	}
	return &r, nil

}

// GetIncidents return specific endpoint data
func (s *Session) GetIncidents() (*IncidentData, error) {

	var r IncidentData
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := s.HTTPService.Get(ctx, s.url(Incidents), &r); err != nil {
		return nil, err
	}
	return &r, nil

}

// GetNotificationChannels return specific endpoint data
func (s *Session) GetNotificationChannels() (*NotificationChannelData, error) {

	var r NotificationChannelData
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := s.HTTPService.Get(ctx, s.url(NotificationChannels), &r); err != nil {
		return nil, err
	}
	return &r, nil

}

// GetNotificationPolicies return specific endpoint data
func (s *Session) GetNotificationPolicies() (*NotificationPolicyData, error) {

	var r NotificationPolicyData
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := s.HTTPService.Get(ctx, s.url(NotificationPolicies), &r); err != nil {
		return nil, err
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
		fmt.Println(Message, s.deleter(i))
	}

}

func (s *Session) deleter(p interface{}) interface{} {
	var r DeleteMsg

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()
	err := s.HTTPService.Delete(ctx, s.url(p.(string)), &r)

	CheckError(err)
	return r.Message

}

// ApplyManifest post manifest to API to create resource
func (s *Session) ApplyManifest(doc TypeMeta) {

	manifest := ManifestReq{
		Manifest: doc,
	}
	buff, _ := json.Marshal(&manifest)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	r, err := s.HTTPService.Post(ctx, s.url(Manifest), bytes.NewBuffer(buff))
	CheckError(err)
	fmt.Println(r.String())
}
