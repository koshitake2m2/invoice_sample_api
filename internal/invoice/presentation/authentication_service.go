package presentation

import (
	"invoice_sample_api/internal/invoice/domain/user"

	"github.com/labstack/echo/v4"
)

type AuthenticationService interface {
	// リクエストからユーザーを認証し, ユーザー情報を返す
	Authenticate(c echo.Context) (*user.User, error)
}
