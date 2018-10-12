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
	"context"
)

type Session struct {
	parent    *Client
	SessionID string
}

func (s *Session) url(endpoint string) string {
	u := s.parent.urlBase(endpoint)
	return u
}

// GetHealthStatus return Health check data
func (s *Session) GetHealthStatus(ctx context.Context) (*HealthData, error) {
	var r HealthData
	if err := s.parent.doReqURL(ctx, s.url("health/summary"), &r); err != nil {
		return nil, err
	}
	return &r, nil
}
