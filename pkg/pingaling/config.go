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
	"encoding/json"
	"io/ioutil"

	"github.com/ghodss/yaml"
	homedir "github.com/mitchellh/go-homedir"
)

// Config is a struct of configuration file data
type Config struct {
	CurrentServer string   `json:"current-server" yaml:"current-server"`
	Servers       []Server `json:"servers" yaml:"servers"`
}

type Server struct {
	URI  string `json:"server" yaml:"server"`
	Name string `json:"name" yaml:"name"`
}

// GetServerURI returns the current serverURI
func (c Config) GetServerURI() string {
	name := c.CurrentServer
	servers := c.Servers
	// TODO: Can use a filter here
	for _, svr := range servers {
		if svr.Name == name {
			return svr.URI
		}
	}
	return ""
}

// NewConfig reads from .pingaling config file, write into Config struct
func NewConfig(cfgFile string, into interface{}) {
	var (
		content []byte
		err     error
	)

	if cfgFile != "" {
		// Use config file from the flag.
		content, err = ioutil.ReadFile(cfgFile)
		CheckError(err)

	} else {
		// Find home directory.
		home, err := homedir.Dir()
		CheckError(err)
		// Search config in home directory with name ".pingaling".
		content, err = ioutil.ReadFile(home + "/.pingaling")
		CheckError(err)
	}

	if toJSON, err := yaml.YAMLToJSON(content); err != nil {
		panic(err)
	} else {
		if err := json.Unmarshal(toJSON, into); err != nil {
			panic(err)
		}
	}
}
