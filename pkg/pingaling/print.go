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
	"strings"
	"text/tabwriter"
	"time"
)

func FormatDate(date string) string {
	outputFormat := "02 Jan 2006 15:04"
	dateTime, _ := time.Parse(time.RFC3339, date)

	return dateTime.Format(outputFormat)
}

func FormatUrl(url string) string {
	if url == "" {
		return "N/A"
	}
	return url
}

func formatHeaders(headers []string) string {
	for index, header := range headers {
		headers[index] = "# " + strings.ToUpper(header)
	}
	return strings.Join(headers, "\t")
}

// Prints a table to the console
// Parameters
//	headers: slice of strings to use as column headers
//	rows: slice of tab-delimited strings to use as row data
func PrintTable(data FormattedData) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 2, '\t', 0)
	defer w.Flush()

	for index, row := range data.Rows {
		if index == 0 {
			fmt.Fprintf(w, "%s\n", formatHeaders(data.Headers))
		}
		fmt.Fprintf(w, "%s\n", row)
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
