package pingaling

import (
	"strings"
)

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

func FormatHealthStatus(healthData []Health) FormattedData {
	headers := []string{
		"Name",
		"Type",
		"Status",
		"URL",
	}

	data := make([]string, 0)
	for _, healthStatus := range healthData {
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
