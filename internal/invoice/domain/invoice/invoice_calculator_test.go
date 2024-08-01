package invoice_test

import (
	"invoice_sample_api/internal/invoice/domain/invoice"
	"testing"
)

func TestCalculateBillingAmount(t *testing.T) {
	actual := invoice.CalculateBillingAmount(10000, 4)
	if actual != 10440 {
		t.Fatal("when paymentAmount is 10000 and feeRate is 4, billing amount should be 10440")
	}
	// TODO: 他のケースについてもテストを追加してください. paymentAmountが0の場合, feeRateが0の場合など
}
