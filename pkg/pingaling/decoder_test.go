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
	"testing"

	"github.com/stretchr/testify/assert"
)

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
