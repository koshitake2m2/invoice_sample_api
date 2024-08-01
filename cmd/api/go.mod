module invoice_sample_api/cmd/api

go 1.22.5

require github.com/labstack/echo/v4 v4.12.0

require (
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace invoice_sample_api/internal/base => ../../internal/base

replace invoice_sample_api/internal/invoice/domain => ../../internal/invoice/domain

replace invoice_sample_api/internal/invoice/usecase => ../../internal/invoice/usecase

replace invoice_sample_api/internal/invoice/presentation => ../../internal/invoice/presentation

replace invoice_sample_api/internal/invoice/infra => ../../internal/invoice/infra
