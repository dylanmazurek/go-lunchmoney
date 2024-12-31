package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/Rhymond/go-money"
	"github.com/dylanmazurek/go-lunchmoney/pkg/utilities/date"
)

type TransactionFilter struct {
	AssetID int `url:"asset_id"`

	TagID           *int64     `url:"tag_id,omitempty"`
	RecurringID     *int64     `url:"recurring_id,omitempty"`
	CategoryID      *int64     `url:"category_id,omitempty"`
	StartDate       *date.Date `url:"start_date,omitempty"`
	EndDate         *date.Date `url:"end_date,omitempty"`
	DebitAsNegative bool       `url:"debit_as_negative"`

	Limit  int `url:"limit"`
	Offset int `url:"offset"`
}

type TransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
	HasMore      bool          `json:"has_more"`
}

func (a *TransactionResponse) UnmarshalJSON(data []byte) error {
	type Alias TransactionResponse
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	return nil
}

type Transaction struct {
	ID json.Number `json:"id,omitempty"`

	Payee        string `json:"payee"`
	OriginalName string `json:"original_name,omitempty"`
	DisplayName  string `json:"display_name,omitempty"`

	AssetID          *json.Number `json:"asset_id,omitempty"`
	AssetDisplayName string       `json:"asset_display_name,omitempty"`
	AssetStatus      *string      `json:"asset_status,omitempty"`

	Notes             *string            `json:"notes,omitempty"`
	CategoryID        *json.Number       `json:"category_id,omitempty"`
	CategoryName      *string            `json:"category_name,omitempty"`
	CategoryGroupID   *json.Number       `json:"category_group_id,omitempty"`
	CategoryGroupName *string            `json:"category_group_name,omitempty"`
	Status            *TransactionStatus `json:"status,omitempty"`
	IsGroup           *bool              `json:"is_group,omitempty"`
	GroupID           *json.Number       `json:"group_id,omitempty"`
	ParentID          *json.Number       `json:"parent_id,omitempty"`
	ExternalID        *string            `json:"external_id,omitempty"`
	RecurringID       *json.Number       `json:"recurring_id,omitempty"`
	IsIncome          *bool              `json:"is_income,omitempty"`

	Tags []string `json:"-"`

	Amount       *money.Money `json:"-"`
	Date         *date.Date   `json:"-"`
	OriginalDate *date.Date   `json:"-"`
}

type TransactionStatus string

const (
	TransactionStatusCleared   TransactionStatus = "cleared"
	TransactionStatusUncleared TransactionStatus = "uncleared"
)

func (t *Transaction) MarshalJSON() ([]byte, error) {
	type Alias Transaction
	newStruct := struct {
		*Alias
		AmountRaw   *string `json:"amount,omitempty"`
		CurrencyRaw *string `json:"currency,omitempty"`

		DateRaw         *string   `json:"date,omitempty"`
		OriginalDateRaw *string   `json:"original_date,omitempty"`
		Tags            *[]string `json:"tags,omitempty"`
	}{
		Alias: (*Alias)(t),
	}

	if t.Amount != nil {
		a := fmt.Sprintf("%.2f", t.Amount.AsMajorUnits())
		newStruct.AmountRaw = &a
	}

	if t.Amount != nil {
		c := strings.ToLower(t.Amount.Currency().Code)
		newStruct.CurrencyRaw = &c
	}

	if t.Date != nil {
		d := t.Date.String()
		newStruct.DateRaw = &d
	}

	if t.Tags != nil {
		newStruct.Tags = &t.Tags
	}

	marshaledJSON, err := json.Marshal(&newStruct)

	return marshaledJSON, err
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	type Alias Transaction
	aux := &struct {
		*Alias
		AmountRaw   string `json:"amount"`
		CurrencyRaw string `json:"currency"`

		DateRaw         string  `json:"date"`
		OriginalDateRaw *string `json:"original_date,omitempty"`
		Tags            []struct {
			Name string `json:"name"`
			ID   int64  `json:"id"`
		} `json:"tags"`
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.AmountRaw != "" {
		amountFloat, err := strconv.ParseFloat(aux.AmountRaw, 64)
		if err != nil {
			return err
		}

		currency := strings.ToUpper(aux.CurrencyRaw)
		amount := money.NewFromFloat(amountFloat, currency)
		t.Amount = amount
	}

	if aux.DateRaw != "" {
		date, err := date.Parse(aux.DateRaw)
		if err != nil {
			return err
		}

		t.Date = &date
	}

	if len(aux.Tags) > 0 {
		for _, tag := range aux.Tags {
			t.Tags = append(t.Tags, tag.Name)
		}
	}

	return nil
}

type InsertRequest struct {
	Transactions      []Transaction `json:"transactions"`
	ApplyRules        bool          `json:"apply_rules,omitempty"`
	SkipDuplicates    bool          `json:"skip_duplicates,omitempty"`
	CheckForRecurring bool          `json:"check_for_recurring,omitempty"`
	DebitAsNegative   bool          `json:"debit_as_negative,omitempty"`
	SkipBalanceUpdate bool          `json:"skip_balance_update,omitempty"`
}

type InsertResponse struct {
	Errors []string `json:"error,omitempty"`

	Ids []int64 `json:"ids,omitempty"`
}

type UpdateRequest struct {
	Transaction       Transaction `json:"transaction"`
	DebitAsNegative   bool        `json:"debit_as_negative,omitempty"`
	SkipBalanceUpdate bool        `json:"skip_balance_update,omitempty"`
}

type UpdateResponse struct {
	Errors []string `json:"error,omitempty"`

	Updated bool `json:"updated,omitempty"`
}
