package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Static(""))
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Drink the pickle juice! IT'S DELICIOUS 🥒🥤")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
