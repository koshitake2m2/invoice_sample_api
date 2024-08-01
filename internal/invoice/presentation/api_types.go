package presentation

import (
	"time"
)

/*
請求書データ作成 リクエスト
*/
type CreateInvoiceRequest struct {
	IssueDate          time.Time `json:"issue_date"`
	PaymentAmount      float64   `json:"payment_amount"`
	Fee                float64   `json:"fee"`
	FeeRate            float64   `json:"fee_rate"`
	ConsumptionTax     float64   `json:"consumption_tax"`
	ConsumptionTaxRate float64   `json:"consumption_tax_rate"`
	PaymentDueDate     time.Time `json:"payment_due_date"`
	Status             string    `json:"status"`
}

/*
請求書データ一覧取得 レスポンス
*/
type InvoiceListResponse struct {
	Invoices []InvoiceResponse `json:"invoices"`
}

type InvoiceResponse struct {
	InvoiceID          int64     `json:"invoice_id"`
	ClientID           int64     `json:"client_id"`
	IssueDate          time.Time `json:"issue_date"`
	PaymentAmount      float64   `json:"payment_amount"`
	Fee                float64   `json:"fee"`
	FeeRate            float64   `json:"fee_rate"`
	ConsumptionTax     float64   `json:"consumption_tax"`
	ConsumptionTaxRate float64   `json:"consumption_tax_rate"`
	BillingAmount      float64   `json:"billing_amount"`
	PaymentDueDate     time.Time `json:"payment_due_date"`
	Status             string    `json:"status"`
}
