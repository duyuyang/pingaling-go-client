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
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServerURI(t *testing.T) {

	cfg := Config{
		CurrentServer: "localhostv1",
		Servers: []Server{
			Server{
				URI:  "http://localhost/api/v1",
				Name: "localhostv1",
			},
			Server{
				URI:  "http://localhost/api/v2",
				Name: "localhostv2",
			},
		},
	}
	actual := cfg.GetServerURI()
	assert.Equal(t, actual, "http://localhost/api/v1")
}

func TestNewConfig(t *testing.T) {
	cfgFile := filepath.Join("testdata", "config.yml")

	var cfgStruct Config

	NewConfig(cfgFile, &cfgStruct)
	assert.Equal(t, "localhost", cfgStruct.CurrentServer)

}
