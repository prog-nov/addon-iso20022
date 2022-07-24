package main

type PaymentOutput struct {
	Credttm          string `json:"credit_date,omitempty"`
	SettlementAmount string `json:"settlement_amount,omitempty"`
	SettlementCurcy  string `json:"settlememt_currency,omitempty"`
	Desc             string `json:"description,omitempty"`
}
