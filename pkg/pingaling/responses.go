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
