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
