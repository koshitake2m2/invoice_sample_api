package invoice

import (
	"invoice_sample_api/internal/base/libs"
)

type InvoiceRepository interface {
	Create(invoice Invoice, io libs.IOContext) error
}
