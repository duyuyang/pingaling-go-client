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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrIter(t *testing.T) {

	chl := StrIter([]string{"foo", "bar"})

	assert.Equal(t, <-chl, "foo")
	assert.Equal(t, <-chl, "bar")
}

func TestMap(t *testing.T) {
	mapper := func(i interface{}) interface{} {
		return strings.ToUpper(i.(string))
	}
	newMap := Map(mapper, New("foo", "bar", "zoo"))
	assert.Equal(t, <-newMap, "FOO")
	assert.Equal(t, <-newMap, "BAR")
	assert.Equal(t, <-newMap, "ZOO")
}
