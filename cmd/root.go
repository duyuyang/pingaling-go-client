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
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	pl "github.com/spf13/pingaling/pkg/pingaline"
	"github.com/spf13/viper"
)

var cfgFile string
var clientCfg interface{}
var session *pl.Session

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pingaling",
	Short: "Monitoring all the things",
	Long: `Pingaling CLI is a tool to setup your monitoring needs. For example:
 Get monitoring endpoins from command line
 Set up monitoring for cronjob from command line
 Set up PagerDuty alerts in CI pipeline`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pingaling)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".pingaling".
		viper.SetConfigType("yaml")
		viper.SetConfigFile(home + "/.pingaling")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("\nRead Config: ", err)
	} else {
		// read the field
		currentServer := viper.GetString("current-server")
		serversI := viper.Get("servers") // interface{}

		clientCfg = pl.NewConfig(currentServer, serversI)
	}

	// initiate the client
	client := pl.Client{
		BaseURL: clientCfg.(*pl.Config).GetServerURI(),
	}

	// Use session to make functtion call
	session, _ = client.CreateSession()

}
