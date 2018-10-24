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

import "testing"

func TestTableHealth(t *testing.T) {

	h := []Health{
		Health{
			URL:     "http://foo",
			Updated: "foo",
			Type:    "foo",
			Status:  "unhealthy",
			Name:    "foo",
		},
		Health{
			URL:     "http://bar",
			Updated: "bar",
			Type:    "bar",
			Status:  "healthy",
			Name:    "bar",
		},
	}

	TableHealth(h)
}

func TestTableEndpoints(t *testing.T) {
	ep := Endpoint{
		URL:         "http://foo",
		NextCheck:   "foo",
		Name:        "foo",
		Description: "foo",
	}
	TableEndpoints(ep)
}

func TestTableIncidents(t *testing.T) {
	ins := []Incident{
		Incident{
			URL:         "http://localhost/foo",
			UpdatedAt:   "foo",
			Status:      "foo",
			NextAttempt: "foo",
			Name:        "foo",
			ID:          3,
		},
	}
	TableIncidents(ins)
}

func TestTableNotificationChannels(t *testing.T) {
	ncs := []NotificationChannel{
		NotificationChannel{
			UpdatedAt: "foo",
			Type:      "foo",
			Name:      "foo",
		},
	}

	TableNotificationChannels(ncs)
}

func TestTableNotificationPolicies(t *testing.T) {
	nps := []NotificationPolicy{
		NotificationPolicy{
			UpdatedAt: "foo",
			Type:      "foo",
			Name:      "foo",
			Endpoint:  "foo",
			Channel:   "foo",
		},
	}
	TableNotificationPolicies(nps)
}
