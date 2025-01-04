package lunchmoney

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/constants"
	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
)

type ListTransactionFilter struct {
	AssetID         *int64    `url:"asset_id"`
	Status          *string   `url:"status"`
	StartDate       time.Time `url:"start_date"`
	EndDate         time.Time `url:"end_date"`
	Limit           *int      `url:"limit"`
	IncludePending  bool      `url:"include_pending"`
	DebitAsNegative bool      `url:"debit_as_negative"`
	TagID           *int64    `url:"tag_id"`
}

func (c *Client) ListTransaction(filter ListTransactionFilter) (*[]models.Transaction, error) {
	params := url.Values{}
	if filter.AssetID != nil {
		params.Set("asset_id", fmt.Sprintf("%d", *filter.AssetID))
	}

	if filter.Status != nil {
		params.Set("status", *filter.Status)
	}

	if filter.TagID != nil {
		params.Set("tag_id", fmt.Sprintf("%d", *filter.TagID))
	}

	params.Set("start_date", filter.StartDate.Format("2006-01-02"))

	if !filter.EndDate.IsZero() {
		params.Set("end_date", filter.EndDate.Format("2006-01-02"))
	}

	params.Set("debit_as_negative", fmt.Sprintf("%t", filter.DebitAsNegative))

	if filter.Limit != nil {
		params.Set("limit", fmt.Sprintf("%d", *filter.Limit))
	}

	hasMore := true
	offset := 0

	var transactions []models.Transaction
	for hasMore {
		params.Set("offset", fmt.Sprintf("%d", offset))
		req, err := c.NewRequest(http.MethodGet, constants.API_PATH_TRANSACTIONS, nil, &params)
		if err != nil {
			return nil, err
		}

		var response models.TransactionResponse
		err = c.Do(req, &response)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, response.Transactions...)
		offset += len(response.Transactions)
		hasMore = response.HasMore
		if filter.Limit != nil && offset >= *filter.Limit {
			hasMore = false
		}
	}

	return &transactions, nil
}

func (c *Client) InsertTransactions(transactions []models.Transaction, debitAsNegative bool) (*[]int64, error) {
	insertReqBody := &models.InsertRequest{
		Transactions:      transactions,
		ApplyRules:        true,
		SkipDuplicates:    true,
		CheckForRecurring: true,
		DebitAsNegative:   debitAsNegative,
		SkipBalanceUpdate: true,
	}

	insertJson, err := json.Marshal(&insertReqBody)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest(http.MethodPost, constants.API_PATH_TRANSACTIONS, bytes.NewReader(insertJson), nil)
	if err != nil {
		return nil, err
	}

	var transactionInsertResponse models.InsertResponse
	err = c.Do(req, &transactionInsertResponse)

	if len(transactionInsertResponse.Errors) > 0 {
		for _, oneErr := range transactionInsertResponse.Errors {
			newError := errors.New(oneErr)
			err = errors.Join(err, newError)
		}
	}

	return &transactionInsertResponse.Ids, err
}

func (c *Client) UpdateTransaction(transaction models.Transaction, debitAsNegative bool) (*bool, error) {
	updateReqBody := &models.UpdateRequest{
		Transaction:       transaction,
		DebitAsNegative:   debitAsNegative,
		SkipBalanceUpdate: true,
	}

	updateJson, err := json.Marshal(&updateReqBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", constants.API_PATH_TRANSACTIONS, transaction.ID)
	req, err := c.NewRequest(http.MethodPut, url, bytes.NewReader(updateJson), nil)
	if err != nil {
		return nil, err
	}

	var transactionUpdateResponse models.UpdateResponse
	err = c.Do(req, &transactionUpdateResponse)

	if len(transactionUpdateResponse.Errors) > 0 {
		for _, oneErr := range transactionUpdateResponse.Errors {
			newError := errors.New(oneErr)
			err = errors.Join(err, newError)
		}
	}

	return &transactionUpdateResponse.Updated, err
}
