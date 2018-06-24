package main

import (
	"github.com/h00s/lotoweb/loto"
	"github.com/h00s/lotoweb/templates"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	r := echo.New()
	r.Renderer = templates.LoadTemplates("views/*.html")
	r.Use(middleware.Static("/static"))
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	lc := loto.NewController()
	r.GET("/", lc.Index)

	r.Logger.Fatal(r.Start("localhost:1323"))
}
