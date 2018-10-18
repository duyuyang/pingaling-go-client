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

type createSessionResp struct {
	RetMsg    string `json:"ret_msg"`
	SessionID string `json:"session_id"`
	Timestamp string `json:"timestamp"`
}

type Health struct {
	URL     string `json:"url"`
	Updated string `json:"updated"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	Name    string `json:"name"`
}

type HealthData struct {
	Data []Health `json:"data"`
}

type Endpoint struct {
	URL         string `json:"url"`
	NextCheck   string `json:"next_check"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EndpointData struct {
	Data Endpoint `json:"data"`
}

type EndpointManifestReq struct {
	Manifest   EndpointMReqSpec `json:"spec"`
	Kind       string           `json:"kind"`
	APIVersion string           `json:"apiVersion"`
}

type EndpointMReqSpec struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type EndpointManifestResp struct {
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Incident struct {
	URL         string `json:"url"`
	UpdatedAt   string `json:"updated_at"`
	Status      string `json:"status"`
	NextAttempt string `json:"next_attempt"`
	Name        string `json:"name"`
	ID          int    `json:"id"`
}

type IncidentData struct {
	Data []Incident `json:"data"`
}

type NotificationChannel struct {
	UpdatedAt string `json:"updated_at"`
	Type      string `json:"type"`
	Name      string `json:"name"`
}

type NotificationChannelData struct {
	Data []NotificationChannel `json:"data"`
}

type NotificationPolicy struct {
	UpdatedAt string `json:"updated_at"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Endpoint  string `json:"endpoint"`
	Channel   string `json:"channel"`
}

type NotificationPolicyData struct {
	Data []NotificationPolicy `json:"data"`
}

type DeleteMsg struct {
	Message string `json:"message"`
}

type TypeMeta struct {
	APIVersion int                    `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	Kind       string                 `json:"kind,omitempty" yaml:"kind,omitempty"`
	Spec       map[string]interface{} `json:"spec" yaml:"spec"`
}

type ManifestReq struct {
	Manifest TypeMeta `json:"manifest"`
}
