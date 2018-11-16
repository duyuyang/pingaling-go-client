package pingaling

import "strings"

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
