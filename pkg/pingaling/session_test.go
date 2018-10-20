package pingaling

import (
	"bytes"
	"encoding/json"
	"reflect"
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

func TestSessionsTestSuite(t *testing.T) {
	suite.Run(t, new(SessionsTestSuite))
}

func (suite *SessionsTestSuite) TestURL() {

	assert.Equal(suite.T(), suite.mockURL+"/blah", suite.mockSession.url("blah"))
}

func (suite *SessionsTestSuite) TestGetHealthStatus() {

	mockResp := bytes.NewBuffer([]byte(`
	{
		"data": [
			{
				"url": "https://service.svc.local/healthz",
				"updated": "2018-10-16T20:46:34.736998Z",
				"type": "endpoint",
				"status": "pending",
				"name": "my-service24"
			},
			{
				"url": "http://foobar.com.au/diagnostic",
				"updated": "2018-10-16T20:46:34.736362Z",
				"type": "endpoint",
				"status": "pending",
				"name": "my-service25"
			}
		]
	}`))

	suite.clt.On("Get", suite.mockURL+"/health/summary").Return(*mockResp, nil)

	resp, err := suite.mockSession.GetHealthStatus()
	assert.Equal(suite.T(), reflect.TypeOf(resp).String(), "*pingaling.HealthData")
	assert.NotEmpty(suite.T(), resp.Data)
	assert.Nil(suite.T(), err)

}

func (suite *SessionsTestSuite) TestDeleter() {

	mockResp := bytes.NewBuffer([]byte(`{"Message": "Test Delete Message"}`))
	suite.clt.On("Delete", suite.mockURL+"/test").Return(*mockResp, nil)

	resp := suite.mockSession.deleter("test")
	assert.Equal(suite.T(), "Test Delete Message", resp.(string))
}

func (suite *SessionsTestSuite) TestApplyManifest() {

	docData := bytes.NewBuffer([]byte(`
	{		
		"spec": {
			"name": "periodic-yak-shaver",
			"alert_without_success": {
				"minutes": 3
			}
		},
		"kind": "checks/cronjob",
		"apiVersion": 1
	}`))

	mockResp := bytes.NewBuffer([]byte(`
	{
		"name": "periodic-yak-shaver",
		"description": null
	}`))

	var doc TypeMeta
	JSONDecoder(*docData, doc)
	manifest := ManifestReq{
		Manifest: doc,
	}
	buff, _ := json.Marshal(&manifest)

	suite.clt.On("Post", suite.mockURL+"/manifest", bytes.NewBuffer(buff)).Return(*mockResp, nil)
	b, err := suite.mockSession.ApplyManifest(doc)
	assert.Equal(suite.T(), b.String(), "\n\t{\n\t\t\"name\": \"periodic-yak-shaver\",\n\t\t\"description\": null\n\t}")
	assert.Nil(suite.T(), err)

}
