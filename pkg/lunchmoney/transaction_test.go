package lunchmoney

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
	"github.com/stretchr/testify/assert"
)

func TestListTransaction(t *testing.T) {
	client := &Client{
		internalClient: &http.Client{},
	}

	expectedResponse := &models.TransactionResponse{
		Transactions: []models.Transaction{
			{
				ID:            json.Number("1"),
				Payee:         "Test Payee",
				OriginalName:  "Test Original Name",
				DisplayName:   "Test Display Name",
				AssetID:       json.Number("1"),
				AssetStatus:   stringPtr("active"),
				CategoryID:    json.Number("1"),
				CategoryName:  stringPtr("Test Category"),
				Status:        statusPtr(models.TransactionStatusCleared),
				IsGroup:       boolPtr(false),
				Amount:        moneyPtr(100.0, "USD"),
				Date:          datePtr("2023-01-01"),
				OriginalDate:  datePtr("2023-01-01"),
				Tags:          []string{"tag1", "tag2"},
			},
		},
		HasMore: false,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	filter := ListTransactionFilter{
		StartDate: time.Now().AddDate(0, 0, -30),
		EndDate:   time.Now(),
	}

	transactions, err := client.ListTransaction(filter)
	assert.NoError(t, err)
	assert.NotNil(t, transactions)
	assert.Equal(t, expectedResponse.Transactions, *transactions)
}

func TestInsertTransactions(t *testing.T) {
	client := &Client{
		internalClient: &http.Client{},
	}

	expectedResponse := &models.InsertResponse{
		Ids: []int64{1, 2, 3},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	transactions := []models.Transaction{
		{
			Payee:         "Test Payee",
			OriginalName:  "Test Original Name",
			DisplayName:   "Test Display Name",
			AssetID:       json.Number("1"),
			AssetStatus:   stringPtr("active"),
			CategoryID:    json.Number("1"),
			CategoryName:  stringPtr("Test Category"),
			Status:        statusPtr(models.TransactionStatusCleared),
			IsGroup:       boolPtr(false),
			Amount:        moneyPtr(100.0, "USD"),
			Date:          datePtr("2023-01-01"),
			OriginalDate:  datePtr("2023-01-01"),
			Tags:          []string{"tag1", "tag2"},
		},
	}

	ids, err := client.InsertTransactions(transactions, true)
	assert.NoError(t, err)
	assert.NotNil(t, ids)
	assert.Equal(t, expectedResponse.Ids, *ids)
}

func TestUpdateTransaction(t *testing.T) {
	client := &Client{
		internalClient: &http.Client{},
	}

	expectedResponse := &models.UpdateResponse{
		Updated: true,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	transaction := models.Transaction{
		ID:            json.Number("1"),
		Payee:         "Test Payee",
		OriginalName:  "Test Original Name",
		DisplayName:   "Test Display Name",
		AssetID:       json.Number("1"),
		AssetStatus:   stringPtr("active"),
		CategoryID:    json.Number("1"),
		CategoryName:  stringPtr("Test Category"),
		Status:        statusPtr(models.TransactionStatusCleared),
		IsGroup:       boolPtr(false),
		Amount:        moneyPtr(100.0, "USD"),
		Date:          datePtr("2023-01-01"),
		OriginalDate:  datePtr("2023-01-01"),
		Tags:          []string{"tag1", "tag2"},
	}

	updated, err := client.UpdateTransaction(transaction, true)
	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, expectedResponse.Updated, *updated)
}

func stringPtr(s string) *string {
	return &s
}

func statusPtr(s models.TransactionStatus) *models.TransactionStatus {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func moneyPtr(amount float64, currency string) *money.Money {
	return money.NewFromFloat(amount, currency)
}

func datePtr(dateStr string) *date.Date {
	date, _ := date.Parse(dateStr)
	return &date
}
