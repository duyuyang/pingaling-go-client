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

// Config is a struct of configuration file data
type Config struct {
	currentServer string
	servers       []server
}

type server struct {
	uri  string
	name string
}

// NewConfig returns an initialized Config instance.
func NewConfig(currentServer string, serversI interface{}) *Config {
	return &Config{
		currentServer: currentServer,
		servers:       serverParser(serversI), // []server
	}
}

func serverParser(serversI interface{}) []server {

	servers := make([]server, 0)
	serversS := serversI.([]interface{}) // []interface{}

	for _, svr := range serversS {
		// server is an interface{}

		serverMap := svr.(map[interface{}]interface{})
		// serverMap is a map[interface {}]interface {}

		var s server
		s.name = serverMap["name"].(string) // interface type assertion
		s.uri = serverMap["server"].(string)

		servers = append(servers, s)
	}
	return servers
}

// GetServerURI returns the current serverURI
func (c *Config) GetServerURI() string {
	name := c.currentServer
	servers := c.servers
	// TODO: Can use a filter here
	for _, svr := range servers {
		if svr.name == name {
			return svr.uri
		}
	}
	return ""
}
