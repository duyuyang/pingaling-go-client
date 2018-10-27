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
	"io/ioutil"
	"log"

	pl "bitbucket.org/pingaling-monitoring/client/pkg/pingaling"
	"github.com/spf13/cobra"
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

		// Readfile
		if content, err := ioutil.ReadFile(filename); err != nil {
			log.Fatalf("failed to read manifest: %v, %v", filename, err)
		} else {
			// Split the YAML base on ---
			if docs, err := pl.SplitYAMLDocuments(content); err != nil {
				log.Fatalf("failed to split manifests: %v, %v", content, err)
			} else {
				// Post manifest to API
				for _, d := range docs {
					if buf, err := session.ApplyManifest(d); err != nil {
						log.Printf("failed to create resource %v", err)
					} else {
						fmt.Println(buf.String())
					}
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	applyCmd.Flags().StringVarP(&filename, "filename", "f", "", "File that contains configuratino to apply")

}
