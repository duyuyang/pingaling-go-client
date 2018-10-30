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
	"testing"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func (f Fake) JSONUnmarshal(data []byte, v interface{}) error {
	return errors.New("blah")
}

func (f Fake) YAMLUnmarshal(y []byte, o interface{}) error {
	return errors.New("blah")
}

func (f Fake) YAMLToJSON(b []byte) ([]byte, error) {
	return nil, errors.New("blah")
}

func TestSplitYAMLDocuments(t *testing.T) {
	b := []byte(`---
apiVersion: 1
kind: foo
spec:
  name: bar
---
apiVersion: 1
kind: zoo
spec:
  name: cat`)
	tm, err := SplitYAMLDocuments(b)
	assert.Nil(t, err)
	for _, i := range tm {
		assert.Equal(t, 1, i.APIVersion)
	}
}

func TestSplitYAMLDocumentsUnmarshalError(t *testing.T) {
	b := []byte(`---
apiVersion: 1
kind: foo
spec:
  name: bar
---
apiVersion: 1
kind: zoo
spec:
	name: cat`)
	fake := Fake{}
	decoYAMLUnmarshal = fake.YAMLUnmarshal
	_, err := SplitYAMLDocuments(b)
	assert.EqualError(t, errors.Cause(err), "Unexpected YAML: blah from ---\napiVersion: 1\nkind: foo\nspec:\n  name: bar\n")

}

func TestYAMLDecoder(t *testing.T) {

	b := []byte(`---
  name: bar`)
	type into struct {
		Name string `yaml:"name" json:"name"`
	}
	var r into
	err := YAMLDecoder(b, &r)
	assert.Nil(t, err)
	assert.Equal(t, "bar", r.Name)
}

func TestYAMLDecodertoJSONError(t *testing.T) {
	fake := Fake{}
	decoYAMLtoJSON = fake.YAMLToJSON
	b := []byte(``)
	type into struct {
		Name string `yaml:"name" json:"name"`
	}
	var r into
	err := YAMLDecoder(b, &r)
	assert.EqualError(t, err, "Unexpected YAML: blah from ")

	//clean up
	decoYAMLtoJSON = yaml.YAMLToJSON
}

func TestYAMLDecoderUnmarshalError(t *testing.T) {
	fake := Fake{}
	decoYAMLtoJSON = func([]byte) ([]byte, error) {
		return []byte(``), nil
	}
	decoJSONUnmarshal = fake.JSONUnmarshal
	b := []byte(``)
	var r struct {
		Name string `yaml:"name" json:"name"`
	}
	err := YAMLDecoder(b, &r)
	assert.EqualError(t, err, "Unexpected JSON: blah from ")
	// clean up
	decoYAMLtoJSON = yaml.YAMLToJSON
	decoJSONUnmarshal = json.Unmarshal
}
