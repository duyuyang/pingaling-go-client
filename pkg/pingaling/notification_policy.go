package pingaling

import "strings"

func FormatNotificationPolicy(notificationPolicies []NotificationPolicy) FormattedData {
	headers := []string{
		"Name",
		"Type",
		"Endpoint",
		"Channel",
		"Updated At",
	}

	data := make([]string, 0)

	for _, policy := range notificationPolicies {
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
