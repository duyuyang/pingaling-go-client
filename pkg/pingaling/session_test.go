package pingaling

import (
	"bytes"
	"testing"

	"bitbucket.org/pingaling-monitoring/client/pkg/pingaling/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const MockURL = "http://localhost/api"

type SessionsTestSuite struct {
	suite.Suite
	clt         *mocks.HTTPService
	mockSession Session
	mockURL     string
}

func (suite *SessionsTestSuite) SetupTest() {
	suite.clt = new(mocks.HTTPService)
	suite.mockSession = Session{
		parent: &Client{
			BaseURL: MockURL,
		},
		HTTPService: suite.clt,
	}
	suite.mockURL = "http://localhost/api"
}

func (suite *SessionsTestSuite) TestURL(t *testing.T) {

	assert.Equal(t, suite.mockURL+"/blah", suite.mockSession.url("blah"))
}

// func TestGetHealthStatus(t *testing.T) {
// 	clt := new(ClientMock)
// 	mockSession := Session{
// 		parent: &Client{
// 			BaseURL: MockURL,
// 		},
// 		HTTPService: clt,
// 	}

// 	clt.On("Get", context.Background(), MockURL).Return(nil)

// 	resp, err := mockSession.GetHealthStatus()

// 	assert.Nil(t, err)
// 	assert.Equal(t, reflect.TypeOf(resp).String(), "*pingaling.HealthData")
// }

func (suite *SessionsTestSuite) TestDeleter(t *testing.T) {

	mockResp := bytes.NewBuffer([]byte(`{"Message": "Test Delete Message"}`))
	suite.clt.On("Delete", suite.mockURL).Return(*mockResp, nil)

	resp := suite.mockSession.deleter("test")
	assert.Equal(t, "Test Delete Message", resp.(string))
}

func TestApplyManifest(t *testing.T) {

}
