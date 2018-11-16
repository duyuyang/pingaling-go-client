package pingaling

import (
	"fmt"
)

// Endpoint
type Endpoint struct {
	URL         string `json:"url"`
	NextCheck   string `json:"next_check"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// EndpointData list of healthcheck endpoint
type EndpointData struct {
	Data Endpoint `json:"data"`
}

func (endpoint Endpoint) FormatShow() FormattedData {
	headers := []string{
		"Name",
		"Next check",
		"Url",
		"Description",
	}

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
