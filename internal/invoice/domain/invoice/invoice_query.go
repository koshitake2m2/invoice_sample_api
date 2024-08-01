package invoice

import (
	"invoice_sample_api/internal/base/libs"
	"time"
)

type InvoiceQuery interface {
	Search(startDate time.Time, endDate time.Time, io libs.IOContext) ([]Invoice, error)
}
