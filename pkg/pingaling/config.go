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
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// Config is a struct of configuration file data
type Config struct {
	CurrentServer string   `json:"current-server" yaml:"current-server"`
	Servers       []Server `json:"servers" yaml:"servers"`
}

// Server is the API endpoint
type Server struct {
	URI  string `json:"server" yaml:"server"`
	Name string `json:"name" yaml:"name"`
}

// GetServerURI returns the current serverURI
func (c *Config) GetServerURI() string {
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

// define External functions
var ioReadFile = ioutil.ReadFile
var cfgYAMLDecoder = YAMLDecoder

// NewConfig reads from .pingaling config file, write into Config struct
func (c *Config) NewConfig(cfgFile string) error {
	var (
		content []byte
		err     error
	)

	if cfgFile != "" {
		// Use config file from the flag.
		content, err = ioReadFile(cfgFile)
		if err != nil {
			return errors.Wrapf(err, "failed to read config: %v", cfgFile)
		}

	} else {
		// Find home directory.
		home := os.Getenv("HOME")

		// Search config in home directory with name ".pingaling".
		content, err = ioReadFile(home + "/.pingaling")
		if err != nil {
			return errors.Wrapf(err, "failed to read config %v", home+"/.pingaling")
		}
	}

	if err = cfgYAMLDecoder(content, c); err != nil {
		return errors.Wrap(err, "failed to decode config")
	}

	return nil
}
