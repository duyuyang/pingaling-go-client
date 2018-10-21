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
