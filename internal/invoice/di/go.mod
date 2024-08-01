module invoice_sample_api/internal/invoice/di

go 1.22.5

require (
	github.com/google/subcommands v1.2.0 // indirect
	github.com/google/wire v0.6.0 // indirect
	github.com/labstack/echo/v4 v4.12.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
)

replace invoice_sample_api/internal/base => ../../base
replace invoice_sample_api/internal/invoice/domain => ../domain
replace invoice_sample_api/internal/invoice/usecase => ../usecase
replace invoice_sample_api/internal/invoice/presentation => ../presentation
replace invoice_sample_api/internal/invoice/infra => ../infra
