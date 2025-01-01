package lunchmoney

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestNewRequest(t *testing.T) {
	client := &Client{
		internalClient: &http.Client{},
	}

	req, err := client.NewRequest(http.MethodGet, "/test", nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)
	assert.Equal(t, http.MethodGet, req.HTTPRequest.Method)
	assert.Equal(t, "/test", req.HTTPRequest.URL.Path)
}

func TestDo(t *testing.T) {
	client := &Client{
		internalClient: &http.Client{},
	}

	expectedResponse := &models.Asset{
		AssetID: int64Ptr(1),
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	req, err := client.NewRequest(http.MethodGet, server.URL, nil, nil)
	assert.NoError(t, err)

	var response models.Asset
	err = client.Do(req, &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, &response)
}

func int64Ptr(i int64) *int64 {
	return &i
}
