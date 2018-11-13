// Copyright © 2018 The Pingaling Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pingaling

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

func TableServers(servers []Server) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 2, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "%s\t%s\t%s\n", "# CURRENT", "# NAME", "# URL")

	for _, server := range servers {
		fmt.Fprintf(w, "%v\t%v\t%v\n", boolToString(server.Current), server.Name, server.URI)
	}
}

func boolToString(value bool) string {
	if value {
		return "*"
	} else {
		return ""
	}
}
