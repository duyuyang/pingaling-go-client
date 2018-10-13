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

package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete resources by name",
	Example: `
  # Delete an incident
  pingaling delete incident foo
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

var deleteEndpointCmd = &cobra.Command{
	Use:   "endpoint",
	Short: "Delete the endpoint by name",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires one endpoint resource")
		}
		return nil
	},
	Example: `
 # Delete the endpoint
 pingaling delete endpoint foo
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ep, err := session.DeleteEndpoint(args[0])
		checkError(err)
		fmt.Println("Message: ", ep.Message)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteEndpointCmd)
}
