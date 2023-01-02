package main

import (
	"log"

	"github.com/h00s/lotoweb/app/controllers"
	"github.com/h00s/lotoweb/app/views"
	"github.com/h00s/lotoweb/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	c, err := config.Load("configuration.json")
	if err != nil {
		log.Fatal(err)
	}

	r := echo.New()
	r.Renderer = views.LoadViews(c.Views)
	r.Use(middleware.Static("assets"))
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	lc := controllers.NewLotoController()
	r.GET("/6od45", lc.Loto6od45)
	r.GET("/7od39", lc.Loto7od39)
	r.GET("/", lc.Index)

	r.Logger.Fatal(r.Start("0.0.0.0:8080"))
}
