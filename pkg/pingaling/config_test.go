package pingaling

import (
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
