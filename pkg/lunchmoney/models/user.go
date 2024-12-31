package models

type User struct {
	ID          int     `json:"user_id"`
	Name        string  `json:"user_name"`
	Email       string  `json:"user_email"`
	AccountID   int     `json:"account_id"`
	BudgetName  *string `json:"budget_name"`
	APIKeyLabel *string `json:"api_key_label"`
}
