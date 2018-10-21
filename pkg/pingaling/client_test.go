package pingaling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	var client Client
	var s interface{}
	s = &Session{}
	session, err := client.CreateSession()
	assert.Nil(t, err)
	assert.IsType(t, s, session)
}

func TestURLBase(t *testing.T) {
	var client Client
	URL := client.urlBase("foo")
	assert.Equal(t, "http://localhost:4000/api/foo", URL)
}
