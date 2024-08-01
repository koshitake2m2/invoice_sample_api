package usecase

import (
	"time"
)

/*
請求書データ作成 ユースケースの入力
*/
type CreateInvoiceInput struct {
	IssueDate          time.Time
	ClientID           int64
	PaymentAmount      float64
	Fee                float64
	FeeRate            float64
	ConsumptionTax     float64
	ConsumptionTaxRate float64
	PaymentDueDate     time.Time
	Status             string
}

/*
請求書データ一覧取得 ユースケースの入力
*/
type InvoiceListInput struct {
	StartDate time.Time
	EndDate   time.Time
}

/*
請求書データ一覧取得 ユースケースの出力
*/
type InvoiceListOutput struct {
	Invoices []InvoiceOutput
}

type InvoiceOutput struct {
	InvoiceID          int64
	ClientID           int64
	IssueDate          time.Time
	PaymentAmount      float64
	Fee                float64
	FeeRate            float64
	ConsumptionTax     float64
	ConsumptionTaxRate float64
	BillingAmount      float64
	PaymentDueDate     time.Time
	Status             string
}
