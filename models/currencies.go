package models

// CurrencyModel represents a currency model
type Currency struct {
	Code   string `json:"code"`   // The code of the currency
	Name   string `json:"name"`   // The name of the currency
	Symbol string `json:"symbol"` // The symbol of the currency
}
