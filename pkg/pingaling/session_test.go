package pingaling

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MockURL = "http://locathost/api"

func TestURL(t *testing.T) {
	clt := new(ClientMock)
	mockSession := Session{
		parent: &Client{
			BaseURL: MockURL,
		},
		HTTPService: clt,
	}
	assert.Equal(t, MockURL+"/blah", mockSession.url("blah"))
}

func TestGetHealthStatus(t *testing.T) {
	clt := new(ClientMock)
	mockSession := Session{
		parent: &Client{
			BaseURL: MockURL,
		},
		HTTPService: clt,
	}

	clt.On("Get", context.Background(), MockURL).Return(nil)

	resp, err := mockSession.GetHealthStatus()

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(resp).String(), "*pingaling.HealthData")
}

func TestDeleter(t *testing.T) {
	clt := new(ClientMock)
	mockSession := Session{
		parent: &Client{
			BaseURL: MockURL,
		},
		HTTPService: clt,
	}

	var mockI interface{}

	mockI = "Test"
	clt.On("Delete", context.Background(), MockURL).Return(nil)

	resp := mockSession.deleter(mockI)
	assert.Equal(t, "Test Delete Message", resp.(string))
}

func TestApplyManifest(t *testing.T){
	
}