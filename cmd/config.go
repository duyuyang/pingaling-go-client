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
	"bitbucket.org/pingaling-monitoring/client/pkg/pingaling"
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Modify $HOME/.pingaling file",
	Long:  `Use config subcommand to switch between backend servers`,
	Example: `
 # Get the current server
 pingaling config cs
`,
}

// currentServerCmd represents the currentServer command
var currentServerCmd = &cobra.Command{
	Use:     "current-server",
	Aliases: []string{"cs"},
	Short:   "Show the current backend API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("current Server: ", cfgStruct.GetServerURI())
	},
}

var serversCmd = &cobra.Command{
	Use:     "servers",
	Short:   "List all servers",
	Run: func(cmd *cobra.Command, args []string) {
		pingaling.PrintTable(cfgStruct.ListServers())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(currentServerCmd)
	configCmd.AddCommand(serversCmd)
}
