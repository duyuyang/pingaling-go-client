package pingaline

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// TableHealth format the health data to print
func TableHealth(h []Health) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 2, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n%s\t%s\t%s\t%s", "# Name", "# TYPE", "# STATUS", "# URL")

	for _, v := range h {
		fmt.Fprintf(w, "\n%v\t%v\t%v\t\"%v\"", v.Name, v.Type, v.Status, v.URL)
	}
}

// TableEndpoints format the Endpoints data to print
func TableEndpoints(ep Endpoint) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 2, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n%s\t%s\t%s\t%s", "# Name", "# Next Check", "# URL", "# DESCRIPTION")

	fmt.Fprintf(w, "\n%v\t%v\t%v\t\"%v\"", ep.Name, ep.NextCheck, ep.URL, ep.Description)

}

// TableIncidents format the health data to print
func TableIncidents(ins []Incident) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 2, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n%s\t%s\t%s\t%s\t%s\t%s", "# Name", "# ID", "# STATUS", "# UPDATE AT", "# NEXT ATTEMPT", "# URL")

	for _, v := range ins {
		fmt.Fprintf(w, "\n%v\t%v\t%v\t%v\t%v\t\"%v\"", v.Name, v.ID, v.Status, v.UpdatedAt, v.NextAttempt, v.URL)
	}
}

// TableNotificationChannels format the health data to print
func TableNotificationChannels(ncs []NotificationChannel) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 2, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n%s\t%s\t%s", "# Name", "# TYPE", "# UPDATED AT")

	for _, v := range ncs {
		fmt.Fprintf(w, "\n%v\t%v\t%v", v.Name, v.Type, v.UpdatedAt)
	}
}

// TableNotificationPolicies format the health data to print
func TableNotificationPolicies(nps []NotificationPolicy) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 2, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n%s\t%s\t%s\t%s\t%s", "# Name", "# TYPE", "# ENDPOINT", "# CHANNEL", "# UPDATED AT")

	for _, v := range nps {
		fmt.Fprintf(w, "\n%v\t%v\t%v\t%v\t%v", v.Name, v.Type, v.Endpoint, v.Channel, v.UpdatedAt)
	}
}
