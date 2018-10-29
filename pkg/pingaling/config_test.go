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
	"path/filepath"
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
)

type Fake struct{}

func (f Fake) ReadFile(filename string) ([]byte, error) {
	return []byte{}, errors.New("blah")
}

func (f Fake) YAMLDecoder(b []byte, into interface{}) error {
	return errors.New("blah")
}

func TestGetServerURI(t *testing.T) {

	cfg := &Config{
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

func TestGetServerURINotFound(t *testing.T) {

	cfg := &Config{
		CurrentServer: "localhost",
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
	assert.Equal(t, actual, "")
}

func TestNewConfig(t *testing.T) {
	cfgFile := filepath.Join("testdata", "config.yml")

	var cfgStruct Config

	err := cfgStruct.NewConfig(cfgFile)
	assert.Nil(t, err)
	assert.Equal(t, "localhost", cfgStruct.CurrentServer)

}

func TestNewConfigcfgFileError(t *testing.T) {
	cfgFile := filepath.Join("testdata", "config.yml")
	var cfgStruct Config
	fake := Fake{}
	ioReadFile = fake.ReadFile
	err := cfgStruct.NewConfig(cfgFile)
	assert.NotNil(t, err)
	assert.EqualError(t, errors.Cause(err), "blah")
	// clean up
	ioReadFile = ioutil.ReadFile
}

func TestNewConfigDecoderError(t *testing.T) {
	cfgFile := filepath.Join("testdata", "config.yml")

	var cfgStruct Config
	fake := Fake{}
	cfgYAMLDecoder = fake.YAMLDecoder
	err := cfgStruct.NewConfig(cfgFile)
	assert.Equal(t, "", cfgStruct.CurrentServer)
	assert.EqualError(t, err, "failed to decode config: blah")
	assert.EqualError(t, errors.Cause(err), "blah")
	// clean up
	cfgYAMLDecoder = YAMLDecoder
}

func TestNewConfigHomedir(t *testing.T) {

	var cfgStruct Config
	fake := Fake{}
	ioReadFile = fake.ReadFile
	err := cfgStruct.NewConfig("")
	assert.NotNil(t, err)
	assert.EqualError(t, errors.Cause(err), "blah")
	// clean up
	ioReadFile = ioutil.ReadFile
}
