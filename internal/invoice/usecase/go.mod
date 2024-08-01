module invoice_sample_api/internal/invoice/usecase

go 1.22.5

replace invoice_sample_api/internal/base => ../../base

replace invoice_sample_api/internal/invoice/domain => ../domain

require (
	invoice_sample_api/internal/base v0.0.0-00010101000000-000000000000
	invoice_sample_api/internal/invoice/domain v0.0.0-00010101000000-000000000000
)
