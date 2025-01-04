package lunchmoney

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientIntegration(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	start := time.Now().AddDate(0, 0, -30)
	end := time.Now()

	filter := ListTransactionFilter{
		StartDate: start,
		EndDate:   end,
	}

	transactions, err := client.ListTransaction(filter)
	assert.NoError(t, err)
	assert.NotNil(t, transactions)
	assert.Greater(t, len(*transactions), 0)
}

func TestAuthClientIntegration(t *testing.T) {
	ctx := context.Background()
	opts := Options{}

	authClient, err := NewAuthClient(ctx, opts)
	assert.NoError(t, err)
	assert.NotNil(t, authClient)

	client, err := authClient.InitTransportSession()
	assert.NoError(t, err)
	assert.NotNil(t, client)
}
