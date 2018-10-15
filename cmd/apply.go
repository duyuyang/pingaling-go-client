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
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var filename string

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a configuration to a resource by filename",
	Example: `
	# Apply the configuration in endpoint manifest
  pingaling apply -f endpoint.yml
	`,
	Run: func(cmd *cobra.Command, args []string) {

		// get body from config
		
		viper.AddConfigPath(filepath.Dir(filename))
		base := filepath.Base(filename)
		extension := filepath.Ext(filename)
		viper.SetConfigName(base[0 : len(base)-len(extension)]) // Filename without ext

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s", err))
		} else {
			fmt.Println(viper.Get("apiVersion"))
			fmt.Println(viper.GetString("kind"))
			fmt.Println(viper.GetStringMapString("spec"))

		}

		// Pass body to ApplyManifests
		//session.ApplyManifests()

	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	applyCmd.Flags().StringVarP(&filename, "filename", "f", "", "File that contains configuratino to apply")

}
