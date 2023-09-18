package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/status", HttpStatus)
	e.Logger.Fatal(e.Start(":8080"))
}

func HttpStatus(c echo.Context) error {
	return c.String(http.StatusOK, "(>_<)")
}
