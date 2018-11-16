package pingaling

import (
	"strconv"
	"strings"
)

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
