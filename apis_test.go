package gokong

import (
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ApisGetById(t *testing.T) {
	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		Uris:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamUrl:            "http://localhost:4140/testservice",
		StripUri:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HttpsOnly:              true,
		HttpIfTerminated:       true,
	}

	apiClient := NewClient(NewDefaultConfig()).Apis()
	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	result, err := apiClient.GetById(createdApi.Id)

	assert.NotNil(t, result)
	assert.NotNil(t, result)
	assert.Equal(t, createdApi.Id, result.Id)
	assert.Equal(t, createdApi.CreatedAt, result.CreatedAt)
	assert.Equal(t, createdApi.Name, result.Name)
	assert.Equal(t, createdApi.Hosts, result.Hosts)
	assert.Equal(t, createdApi.Uris, result.Uris)
	assert.Equal(t, createdApi.Methods, result.Methods)
	assert.Equal(t, createdApi.UpstreamUrl, result.UpstreamUrl)
	assert.Equal(t, createdApi.StripUri, result.StripUri)
	assert.Equal(t, createdApi.PreserveHost, result.PreserveHost)
	assert.Equal(t, createdApi.UpstreamConnectTimeout, result.UpstreamConnectTimeout)
	assert.Equal(t, createdApi.UpstreamSendTimeout, result.UpstreamSendTimeout)
	assert.Equal(t, createdApi.UpstreamReadTimeout, result.UpstreamReadTimeout)
	assert.Equal(t, createdApi.HttpsOnly, result.HttpsOnly)
	assert.Equal(t, createdApi.HttpIfTerminated, result.HttpIfTerminated)

}

func Test_ApisGetAll(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"another.com"},
		Uris:                   []string{"/another"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/myservice",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}
	apiClient := NewClient(NewDefaultConfig()).Apis()
	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	results, err := apiClient.GetAll()

	assert.True(t, results.Total > 0)
	assert.True(t, len(results.Results) > 0)

}

func Test_ApisCreate(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		Uris:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamUrl:            "http://localhost:4140/testservice",
		StripUri:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HttpsOnly:              true,
		HttpIfTerminated:       true,
	}

	result, err := NewClient(NewDefaultConfig()).Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, apiRequest.Name, result.Name)
	assert.Equal(t, apiRequest.Hosts, result.Hosts)
	assert.Equal(t, apiRequest.Uris, result.Uris)
	assert.Equal(t, apiRequest.Methods, result.Methods)
	assert.Equal(t, apiRequest.UpstreamUrl, result.UpstreamUrl)
	assert.Equal(t, apiRequest.StripUri, result.StripUri)
	assert.Equal(t, apiRequest.PreserveHost, result.PreserveHost)
	assert.Equal(t, apiRequest.UpstreamConnectTimeout, result.UpstreamConnectTimeout)
	assert.Equal(t, apiRequest.UpstreamSendTimeout, result.UpstreamSendTimeout)
	assert.Equal(t, apiRequest.UpstreamReadTimeout, result.UpstreamReadTimeout)
	assert.Equal(t, apiRequest.HttpsOnly, result.HttpsOnly)
	assert.Equal(t, apiRequest.HttpIfTerminated, result.HttpIfTerminated)

}