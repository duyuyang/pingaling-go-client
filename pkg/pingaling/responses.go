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

type CreateSessionResp struct {
	RetMsg    string `json:"ret_msg"`
	SessionID string `json:"session_id"`
	Timestamp string `json:"timestamp"`
}

// DeleteMsg returns the response of delete request
type DeleteMsg struct {
	Message string `json:"message"`
}

// TypeMeta describes the manifest specifications
type TypeMeta struct {
	APIVersion int                    `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	Kind       string                 `json:"kind,omitempty" yaml:"kind,omitempty"`
	Spec       map[string]interface{} `json:"spec" yaml:"spec"`
}

// ManifestReq wrap the manifest specification into a post request body
type ManifestReq struct {
	Manifest TypeMeta `json:"manifest" yaml:"manifest"`
}

// NotificationChannel describes alert toolings
type NotificationChannel struct {
	UpdatedAt string `json:"updated_at"`
	Type      string `json:"type"`
	Name      string `json:"name"`
}

// NotificationChannelData describes list of alert toolings
type NotificationChannelData struct {
	Data []NotificationChannel `json:"data"`
}

// NotificationPolicy describes how alerts notify user
type NotificationPolicy struct {
	UpdatedAt string `json:"updated_at"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Endpoint  string `json:"endpoint"`
	Channel   string `json:"channel"`
}

// NotificationPolicyData describes list of policies
type NotificationPolicyData struct {
	Data []NotificationPolicy `json:"data"`
}

// Incident describes incident data
type Incident struct {
	URL         string `json:"url"`
	UpdatedAt   string `json:"updated_at"`
	Status      string `json:"status"`
	NextAttempt string `json:"next_attempt"`
	Name        string `json:"name"`
	ID          int    `json:"id"`
}

// IncidentData describes list of incidents
type IncidentData struct {
	Data []Incident `json:"data"`
}

// Health status struct
type Health struct {
	URL     string `json:"url"`
	Updated string `json:"updated"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	Name    string `json:"name"`
}

// HealthData list of Health status
type HealthData struct {
	Data []Health `json:"data"`
}

type FormattedData struct {
	Headers []string
	Rows    []string
}

// Cronjob
type Cronjob struct {
	Description string `json:"description"`
	Status      string `json:"status"`
	Name        string `json:"name"`
}

// CronjobsData list of cronjobs
type CronjobsData struct {
	Data []Cronjob `json:"data"`
}

// CronjobData single cronjob
type CronjobData struct {
	Data Cronjob `json:"data"`
}