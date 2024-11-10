package main

import (
	"github.com/a-h/templ"
	"github.com/mattefara/htmx-go-templ/web/template"
	"github.com/labstack/echo/v4"
)

func HTML(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

var count = 0

func main() {
	e := echo.New()

	e.Static("/web/static", "static")

	e.GET("/", func(c echo.Context) error {
		return HTML(c, template.Index(count))
	})

	e.POST("/add", func(c echo.Context) error {
		count = count + 1
		return HTML(c, template.Index(count))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
