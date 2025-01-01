package lunchmoney

import (
	"context"
	"net/http"
	"testing"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
	"github.com/dylanmazurek/go-lunchmoney/pkg/utilities/vault"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthClient(t *testing.T) {
	ctx := context.Background()
	opts := Options{
		vaultClient: &vault.Client{},
	}

	authClient, err := NewAuthClient(ctx, opts)
	assert.NoError(t, err)
	assert.NotNil(t, authClient)
	assert.NotEmpty(t, authClient.secrets.apiKey)
}

func TestInitTransportSession(t *testing.T) {
	ctx := context.Background()
	opts := Options{
		vaultClient: &vault.Client{},
	}

	authClient, err := NewAuthClient(ctx, opts)
	assert.NoError(t, err)
	assert.NotNil(t, authClient)

	client, err := authClient.InitTransportSession()
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.IsType(t, &http.Client{}, client)
}

func TestGetUserData(t *testing.T) {
	ctx := context.Background()
	opts := Options{
		vaultClient: &vault.Client{},
	}

	authClient, err := NewAuthClient(ctx, opts)
	assert.NoError(t, err)
	assert.NotNil(t, authClient)

	user, err := authClient.getUserData(authClient.secrets.apiKey)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, &models.User{}, user)
}

func TestNewAuthClientWithAPIKey(t *testing.T) {
	ctx := context.Background()
	apiKey := "test-api-key"

	authClient, err := NewAuthClientWithAPIKey(ctx, apiKey)
	assert.NoError(t, err)
	assert.NotNil(t, authClient)
	assert.Equal(t, apiKey, authClient.session.GetAPIKey())
}
