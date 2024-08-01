package invoice

import (
	"time"
)

// 請求書
type Invoice struct {
	InvoiceID          int64     // 請求書ID
	CompanyID          int64     // 企業ID
	ClientID           int64     // 取引先ID
	IssueDate          time.Time // 発行日
	PaymentAmount      float64   // 支払金額
	Fee                float64   // 手数料
	FeeRate            float64   // 手数料率
	ConsumptionTax     float64   // 消費税
	ConsumptionTaxRate float64   // 消費税率
	BillingAmount      float64   // 請求金額
	PaymentDueDate     time.Time // 支払期日
	// TODO: Statusを enum のような型で表現する
	Status string // ステータス (未処理、処理中、支払い済み、エラー)
}
