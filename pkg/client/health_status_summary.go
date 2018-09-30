package client

type HealthStatus struct {
	Url     string `json:"url"`
	Updated string `json:"updated"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	Name    string `json:"name"`
}

type HealthStatuses struct {
	Data []HealthStatus
}

func HealthStatusSummary() []HealthStatus {
	path := "/health/summary"
	var statuses HealthStatuses

	ApiGet(path, &statuses)

	return statuses.Data
}
