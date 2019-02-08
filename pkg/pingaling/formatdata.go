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
	"fmt"
	"strconv"
	"strings"
)

// FormatList cronjob data
func (cronjobData CronjobsData) FormatList() FormattedData {
	headers := []string{
		"Name",
		"Status",
		"Description",
	}

	data := make([]string, 0)

	for _, cronjob := range cronjobData.Data {
		row := []string{
			cronjob.Name,
			cronjob.Status,
			cronjob.Description,
		}

		data = append(data, strings.Join(row, "\t"))
	}

	return FormattedData{
		Headers: headers,
		Rows:    data,
	}
}

// FormatShow the cronjob Data
func (cronjobData CronjobData) FormatShow() FormattedData {
	headers := []string{
		"Name",
		"Status",
		"Description",
	}

	cronjob := cronjobData.Data

	data := fmt.Sprintf(
		"%s\t%s\t%s",
		cronjob.Name,
		cronjob.Status,
		cronjob.Description,
	)

	return FormattedData{
		Headers: headers,
		Rows:    []string{data},
	}
}

// Endpoint data
type Endpoint struct {
	URL         string `json:"url"`
	NextCheck   string `json:"next_check"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

// EndpointsData list of endpoints with health status
type EndpointsData struct {
	Data []Endpoint `json:"data"`
}

// EndpointData single endpoint with health status
type EndpointData struct {
	Data Endpoint `json:"data"`
}

// FormatList endpoint data
func (endpointData EndpointsData) FormatList() FormattedData {
	headers := []string{
		"Name",
		"Status",
		"Next check",
		"Url",
		"Description",
	}

	data := make([]string, 0)

	for _, endpoint := range endpointData.Data {
		row := []string{
			endpoint.Name,
			endpoint.Status,
			FormatDate(endpoint.NextCheck),
			endpoint.URL,
			endpoint.Description,
		}

		data = append(data, strings.Join(row, "\t"))
	}

	return FormattedData{
		Headers: headers,
		Rows:    data,
	}
}

// FormatShow endpoint data
func (endpointData EndpointData) FormatShow() FormattedData {
	headers := []string{
		"Name",
		"Status",
		"Next check",
		"Url",
		"Description",
	}

	endpoint := endpointData.Data

	data := fmt.Sprintf(
		"%s\t%s\t%s\t%s\t%s",
		endpoint.Name,
		endpoint.Status,
		FormatDate(endpoint.NextCheck),
		endpoint.URL,
		endpoint.Description,
	)

	return FormattedData{
		Headers: headers,
		Rows:    []string{data},
	}
}

// FormatList notification policies format the data
func (notificationPolicies NotificationPolicyData) FormatList() FormattedData {
	headers := []string{
		"Name",
		"Type",
		"Endpoint",
		"Channel",
		"Updated At",
	}

	data := make([]string, 0)

	for _, policy := range notificationPolicies.Data {
		row := []string{
			policy.Name,
			policy.Type,
			policy.Endpoint,
			policy.Channel,
			policy.UpdatedAt,
		}
		data = append(data, strings.Join(row, "\t"))
	}

	return FormattedData{
		Headers: headers,
		Rows:    data,
	}
}

// FormatList notification channels data
func (notificationChannels NotificationChannelData) FormatList() FormattedData {
	headers := []string{
		"Name",
		"Type",
		"Updated At",
	}

	data := make([]string, 0)
	for _, channel := range notificationChannels.Data {
		row := []string{
			channel.Name,
			channel.Type,
			FormatDate(channel.UpdatedAt),
		}
		data = append(data, strings.Join(row, "\t"))
	}

	return FormattedData{
		Headers: headers,
		Rows:    data,
	}
}

// FormatList inicident data
func (incidentData IncidentData) FormatList() FormattedData {
	headers := []string{
		"Id",
		"Name",
		"Status",
		"Updated At",
		"Next Attempt",
		"Url",
	}

	data := make([]string, 0)

	for _, incident := range incidentData.Data {
		row := []string{
			strconv.Itoa(incident.ID),
			incident.Name,
			incident.Status,
			FormatDate(incident.UpdatedAt),
			FormatDate(incident.NextAttempt),
			FormatUrl(incident.URL),
		}

		data = append(data, strings.Join(row, "\t"))
	}

	return FormattedData{
		Headers: headers,
		Rows:    data,
	}
}

// FormatList health status data
func (healthData HealthData) FormatList() FormattedData {
	headers := []string{
		"Name",
		"Type",
		"Status",
		"URL",
	}

	data := make([]string, 0)
	for _, healthStatus := range healthData.Data {
		row := []string{
			healthStatus.Name,
			healthStatus.Type,
			healthStatus.Status,
			FormatUrl(healthStatus.URL),
		}
		data = append(data, strings.Join(row, "\t"))
	}

	return FormattedData{
		Headers: headers,
		Rows:    data,
	}
}
