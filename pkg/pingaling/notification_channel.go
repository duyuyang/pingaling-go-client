package pingaling

import "strings"

// NotificationPolicy describes how alerts notify user
type NotificationPolicy struct {
	UpdatedAt string `json:"updated_at"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Endpoint  string `json:"endpoint"`
	Channel   string `json:"channel"`
}

// NotificationPolicyData describes list of policies
type NotificationPolicyData struct {
	Data []NotificationPolicy `json:"data"`
}

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

func FormatNotificationChannels(notificationChannels []NotificationChannel) FormattedData {
	headers := []string{
		"Name",
		"Type",
		"Updated At",
	}

	data := make([]string, 0)
	for _, channel := range notificationChannels {
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
