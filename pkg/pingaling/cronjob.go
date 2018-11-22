package pingaling

import (
	"fmt"
	"strings"
)

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
