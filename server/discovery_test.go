// Copyright 2013 Google Inc. All Rights Reserved.
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

package server

import (
	"encoding/json"
	"fmt"
	"github.com/rwl/go-endpoints/endpoints"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var apiConfigMap endpoints.ApiDescriptor

func init() {
	json.Unmarshal([]byte(apiConfigJson), &apiConfigMap)
}

func TestGenerateDiscoveryDocRest(t *testing.T) {
	baseUrl := "https://tictactoe.appspot.com/_ah/api/tictactoe/v1/"

	body := map[string]interface{}{"baseUrl": baseUrl}
	bodyJson, _ := json.Marshal(body)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, string(bodyJson))
	}))
	defer ts.Close()
	discoveryProxyHost = ts.URL

	doc, err := generateDiscoveryDoc(&apiConfigMap, "rest")

	assert.NoError(t, err)
	assert.NotEmpty(t, doc)

	var apiConfig map[string]interface{}
	err = json.Unmarshal([]byte(doc), &apiConfig)
	assert.NoError(t, err)
	assert.Equal(t, apiConfig["baseUrl"], baseUrl)
}

func TestGenerateDiscoveryDocRpc(t *testing.T) {
	rpcUrl := "https://tictactoe.appspot.com/_ah/api/rpc"
	body := map[string]interface{}{"rpcUrl": rpcUrl}
	bodyJson, _ := json.Marshal(body)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, string(bodyJson))
	}))
	defer ts.Close()
	discoveryProxyHost = ts.URL

	doc, err := generateDiscoveryDoc(&apiConfigMap, "rpc")

	assert.NoError(t, err)
	assert.NotEmpty(t, doc)

	var apiConfig map[string]interface{}
	err = json.Unmarshal([]byte(doc), &apiConfig)
	assert.NoError(t, err)
	assert.Equal(t, apiConfig["rpcUrl"], rpcUrl)
}

func TestGenerateDiscoveryDocInvalidFormat(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "Error", 400)
	}))
	defer ts.Close()

	discoveryProxyHost = ts.URL

	_, err := generateDiscoveryDoc(&apiConfigMap, "blah")
	assert.Error(t, err)
}

func TestGenerateDiscoveryDocBadApiConfig(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "", 503)
	}))
	defer ts.Close()
	discoveryProxyHost = ts.URL

	bad := &endpoints.ApiDescriptor{Name: "none"}
	doc, err := generateDiscoveryDoc(bad, "rpc")

	assert.Error(t, err)
	assert.Empty(t, doc, "")
}

func TestGetStaticFileExisting(t *testing.T) {
	body := "static file body"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, body)
	}))
	defer ts.Close()
	staticProxyHost = ts.URL

	response, responseBody, err := getStaticFile("/_ah/api/static/proxy.html")

	assert.NoError(t, err)
	assert.Equal(t, response.StatusCode, 200)
	assert.Equal(t, body, responseBody)
}
