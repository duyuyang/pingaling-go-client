// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"log"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates shell completion scripts",
	Long: `To load completion
# On Linux, using bash

## install bash-completion

	yum install bash-completion -y

## Add autocompletion to your profile

	echo "source <(pingaling completion bash)" >> ~/.bashrc

# On macOS, using bash

## If running Bash 3.2

	brew install bash-completion

## or, if running Bash 4.1+

	brew install bash-completion@2

## Generate the bash-completion script

	pingaling completion bash > $(brew --prefix)/etc/bash_completion.d/pingaling

## Add the following to ~/.bashrc

	source $(brew --prefix)/etc/bash_completion
	source $(brew --prefix)/etc/bash_completion.d/pingaling

# Using Zsh, oh-my-zsh

	echo $fpath

## Add the autocompletion script into any of the given path

Example:

	mkdir ~/.oh-my-zsh/functions
	pingaling completion zsh > ~/.oh-my-zsh/functions/_pingaling
`,
}

// bashCmd represents the bash command
var bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generate autocompletion for bash",
	Long: `For more details, consult 
	pingaling completion -h`,
	Run: func(cmd *cobra.Command, args []string) {
		err := rootCmd.GenBashCompletion(os.Stdout)
		if err != nil {
			log.Fatalf("Bash autocompletion %v", err)
		}
	},
}

// zshCmd represents the zsh command
var zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generate autocompletion for zsh",
	Long: `For more details, consult 
	pingaling completion -h`,
	Run: func(cmd *cobra.Command, args []string) {
		err := rootCmd.GenZshCompletion(os.Stdout)
		if err != nil {
			log.Fatalf("Zsh autocompletion %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.AddCommand(bashCmd)
	completionCmd.AddCommand(zshCmd)

}
