package awesome_scrambler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)


func RunApp() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi there")
	})

	e.Logger.Fatal(e.Start(":5000"))
}