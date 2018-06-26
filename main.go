package main

import (
	"github.com/h00s/lotoweb/app/controllers"
	"github.com/h00s/lotoweb/templates"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	r := echo.New()
	r.Renderer = templates.LoadTemplates()
	r.Use(middleware.Static("assets"))
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	lc := controllers.NewLotoController()
	r.GET("/", lc.Index)

	r.Logger.Fatal(r.Start("localhost:1323"))
}
