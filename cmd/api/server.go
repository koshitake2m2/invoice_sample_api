package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func main() {
	e := echo.New()

	// FIXME: diをしたものを利用してエンドポイントを登録してください.

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":8080"))
}
