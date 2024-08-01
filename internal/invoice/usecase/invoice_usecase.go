package usecase

import (
	"invoice_sample_api/internal/base/libs"
	"invoice_sample_api/internal/invoice/domain/invoice"
	"invoice_sample_api/internal/invoice/domain/user"
)

type InvoiceUsecase struct {
	InvoiceRepository invoice.InvoiceRepository
	InvoiceQuery      invoice.InvoiceQuery
}

func (u InvoiceUsecase) Create(input CreateInvoiceInput, loginUser user.User, io libs.IOContext) error {
	// FIXME: 取引先の存在チェックを行ってください. ログインユーザの所属する企業の取引先であることをチェックしてください.
	// FIXME: invoiceの属性の事前条件のvalidationを行ってください.

	billingAmount := invoice.CalculateBillingAmount(input.PaymentAmount, input.FeeRate)

	// FIXME: inputをInvoiceに変換してください.
	var newInvoice invoice.Invoice = invoice.Invoice{BillingAmount: billingAmount}
	u.InvoiceRepository.Create(newInvoice, io)
	return nil
}

func (u InvoiceUsecase) List(input InvoiceListInput, loginUser user.User, io libs.IOContext) (*InvoiceListOutput, error) {
	// FIXME: input.StartDateとinput.EndDateの範囲のvalidationを行ってください.
	invoices, _ := u.InvoiceQuery.Search(input.StartDate, input.EndDate, io)

	// FIXME: invoicesをInvoiceListOutputに変換してください.
	var output InvoiceListOutput = InvoiceListOutput{Invoices: []InvoiceOutput{{InvoiceID: invoices[0].InvoiceID}}}
	return &output, nil
}
