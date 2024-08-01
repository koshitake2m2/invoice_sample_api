package presentation

import (
	"invoice_sample_api/internal/base/libs"
	"invoice_sample_api/internal/invoice/usecase"

	"github.com/labstack/echo/v4"
)

type InvoiceController struct {
	AuthenticationService AuthenticationService
	InvoiceUsecase        usecase.InvoiceUsecase
	IOContextFactory      libs.IOContextFactory
}

func (ic InvoiceController) Create(c echo.Context) error {
	user, _ := ic.AuthenticationService.Authenticate(c)
	if user == nil {
		return c.JSON(401, "Unauthorized")
	}
	_, err := ic.IOContextFactory.Transaction(func(io libs.IOContext) (*libs.IOResult, error) {
		// FIXME: リクエストからinputを生成してください.
		var input usecase.CreateInvoiceInput = usecase.CreateInvoiceInput{}
		err := ic.InvoiceUsecase.Create(input, *user, io)
		return nil, err
	})
	if err != nil {
		// FIXME: エラーハンドリングを実装してください.
		return c.JSON(400, "validation error")
	}
	return c.JSON(200, "OK")
}

func (ic InvoiceController) List(c echo.Context) error {
	user, _ := ic.AuthenticationService.Authenticate(c)
	if user == nil {
		return c.JSON(401, "Unauthorized")
	}
	_, err := ic.IOContextFactory.ReadOnly(func(io libs.IOContext) (*libs.IOResult, error) {
		// FIXME: リクエストからinputを生成してください.
		var input usecase.InvoiceListInput = usecase.InvoiceListInput{}
		output, err := ic.InvoiceUsecase.List(input, *user, io)
		return &libs.IOResult{V: output}, err
	})
	if err != nil {
		// FIXME: エラーハンドリングを実装してください.
		return c.JSON(400, "validation error")
	}

	// FIXME: outputからレスポンスを生成してください.
	var response InvoiceListResponse = InvoiceListResponse{}

	return c.JSON(200, response)
}
