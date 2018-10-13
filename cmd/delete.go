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
	Use:     "endpoint",
	Short:   "Delete the endpoint by name",
	Aliases: []string{"ep"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires one endpoint resource")
		}
		return nil
	},
	Example: `
  # Delete the endpoint
  pingaling delete endpoint foo
  pingaling delete ep foo1 foo2 ...
	`,
	Run: func(cmd *cobra.Command, args []string) {
		session.DeleteEndpoints(args)
	},
}

var deleteNotificationChannelCmd = &cobra.Command{
	Use:     "notification-channel",
	Short:   "Delete the notification channels by name",
	Aliases: []string{"nc"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires at least one notification channel resource")
		}
		return nil
	},
	Example: `
  # Delete the notification channel
  pingaling delete notification-channel foo
  pingaling delete nc foo1 foo2 ...
	`,
	Run: func(cmd *cobra.Command, args []string) {
		session.DeleteNotificationChannels(args)
	},
}

var deleteNotificationPolicyCmd = &cobra.Command{
	Use:     "notification-policy",
	Short:   "Delete the notification-policy by name",
	Aliases: []string{"np"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires at least one notification-policy resource")
		}
		return nil
	},
	Example: `
  # Delete the notification-policy
  pingaling delete notification-policy foo
  pingaling delete np foo1 foo2 ...
	`,
	Run: func(cmd *cobra.Command, args []string) {
		session.DeleteNotificationPolicies(args)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteEndpointCmd)
	deleteCmd.AddCommand(deleteNotificationChannelCmd)
	deleteCmd.AddCommand(deleteNotificationPolicyCmd)
}
