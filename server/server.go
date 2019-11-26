package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "html/dist")

	return e.Start(":8080")
}
