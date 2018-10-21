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

func (suite *SessionsTestSuite) TestGetEndpoints() {

	mockResp := bytes.NewBuffer([]byte(`
	{
		"data": {
			"url": "https://service.svc.local/healthz",
			"next_check": null,
			"name": "my-service21",
			"description": null
		}
	}`))

	suite.clt.On("Get", suite.mockURL+"/endpoints/my-service21").Return(*mockResp, nil)

	resp, err := suite.mockSession.GetEndpoints("my-service21")
	assert.Equal(suite.T(), reflect.TypeOf(resp).String(), "*pingaling.EndpointData")
	assert.NotEmpty(suite.T(), resp.Data)
	assert.Nil(suite.T(), err)

}

func (suite *SessionsTestSuite) TestGetIncidents() {

	mockResp := bytes.NewBuffer([]byte(`
	{
		"data": [
			{
				"url": "https://dingbats.svc.local/boop",
				"updated_at": "2018-10-16T20:46:34.729663Z",
				"status": "open",
				"next_attempt": null,
				"name": "my-service23",
				"id": 1516
			},
			{
				"url": "https://dingbats.svc.local/boop",
				"updated_at": "2018-10-16T20:46:34.730420Z",
				"status": "open",
				"next_attempt": null,
				"name": "my-service23",
				"id": 1517
			},
			{
				"url": "https://dingbats.svc.local/boop",
				"updated_at": "2018-10-16T20:46:34.730946Z",
				"status": "open",
				"next_attempt": null,
				"name": "my-service23",
				"id": 1518
			}
		]
	}`))

	suite.clt.On("Get", suite.mockURL+"/incidents").Return(*mockResp, nil)

	resp, err := suite.mockSession.GetIncidents()
	assert.Equal(suite.T(), reflect.TypeOf(resp).String(), "*pingaling.IncidentData")
	assert.NotEmpty(suite.T(), resp.Data)
	assert.Nil(suite.T(), err)

}

func (suite *SessionsTestSuite) TestGetNotificationChannels() {

	mockResp := bytes.NewBuffer([]byte(`
	{
		"data": [
			{
				"updated_at": "2018-10-16T20:46:34.573605Z",
				"type": "pagerduty",
				"name": "channel6"
			},
			{
				"updated_at": "2018-10-16T20:46:34.577289Z",
				"type": "slack",
				"name": "channel7"
			}
		]
	}`))

	suite.clt.On("Get", suite.mockURL+"/notification_channels").Return(*mockResp, nil)

	resp, err := suite.mockSession.GetNotificationChannels()
	assert.Equal(suite.T(), reflect.TypeOf(resp).String(), "*pingaling.NotificationChannelData")
	assert.NotEmpty(suite.T(), resp.Data)
	assert.Nil(suite.T(), err)

}

func (suite *SessionsTestSuite) TestGetNotificationPolicies() {

	mockResp := bytes.NewBuffer([]byte(`
	{
		"data": [
			{
				"updated_at": "2018-10-16T20:46:34.630255Z",
				"type": "pagerduty",
				"name": "notification_policy12",
				"endpoint": "my-service13",
				"channel": "channel14"
			},
			{
				"updated_at": "2018-10-16T20:46:34.620185Z",
				"type": "slack",
				"name": "notification_policy9",
				"endpoint": "my-service10",
				"channel": "channel11"
			}
		]
	}`))

	suite.clt.On("Get", suite.mockURL+"/notification_policies").Return(*mockResp, nil)

	resp, err := suite.mockSession.GetNotificationPolicies()
	assert.Equal(suite.T(), reflect.TypeOf(resp).String(), "*pingaling.NotificationPolicyData")
	assert.NotEmpty(suite.T(), resp.Data)
	assert.Nil(suite.T(), err)

}

func (suite *SessionsTestSuite) TestDeleter() {

	mockResp := bytes.NewBuffer([]byte(`{"Message": "Test Delete Message"}`))
	suite.clt.On("Delete", suite.mockURL+"/test").Return(*mockResp, nil)

	resp := suite.mockSession.deleter("test")
	assert.Equal(suite.T(), "Test Delete Message", resp.(string))
}

func (suite *SessionsTestSuite) TestDeleteEndpoints() {
	mockResp := bytes.NewBuffer([]byte(`{"Message": "Test Delete Message"}`))
	suite.clt.On("Delete", suite.mockURL+"/endpoints/foo").Return(*mockResp, nil)

	suite.mockSession.DeleteEndpoints([]string{"foo"})
}

func (suite *SessionsTestSuite) TestDeleteNotificationChannels() {
	mockResp := bytes.NewBuffer([]byte(`{"Message": "Test Delete Message"}`))
	suite.clt.On("Delete", suite.mockURL+"/notification_channels/foo").Return(*mockResp, nil)

	suite.mockSession.DeleteNotificationChannels([]string{"foo"})
}

func (suite *SessionsTestSuite) TestDeleteNotificationPolicies() {
	mockResp := bytes.NewBuffer([]byte(`{"Message": "Test Delete Message"}`))
	suite.clt.On("Delete", suite.mockURL+"/notification_policies/foo").Return(*mockResp, nil)

	suite.mockSession.DeleteNotificationPolicies([]string{"foo"})
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
