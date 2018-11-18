package pingaling

import (
	"fmt"
	"strings"
)

// Endpoint
type Endpoint struct {
	URL         string `json:"url"`
	NextCheck   string `json:"next_check"`
	Name        string `json:"name"`
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

func (endpointData EndpointsData) FormatList() FormattedData {
	headers := []string{
		"Name",
		"Next check",
		"Url",
		"Description",
	}

	data := make([]string, 0)

	for _, endpoint := range endpointData.Data {
		row := []string {
			endpoint.Name,
			FormatDate(endpoint.NextCheck),
			endpoint.URL,
			endpoint.Description,
		}

		data = append(data, strings.Join(row, "\t"))
	}

	return FormattedData {
		Headers: headers,
		Rows: data,
	}
}

func (endpointData EndpointData) FormatShow() FormattedData {
	headers := []string{
		"Name",
		"Next check",
		"Url",
		"Description",
	}

	endpoint := endpointData.Data

	data := fmt.Sprintf(
		"%s\t%s\t%s\t%s",
		endpoint.Name,
		FormatDate(endpoint.NextCheck),
		endpoint.URL,
		endpoint.Description,
	)

	return FormattedData{
		Headers: headers,
		Rows:    []string{data},
	}
}
