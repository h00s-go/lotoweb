package controllers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo"
)

// LotoController for Loto methods
type LotoController struct {
}

// NewLotoController creates new loto controller
func NewLotoController() *LotoController {
	return &LotoController{}
}

// Index is first page
func (lc *LotoController) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "loto/index", template.HTML("<p>HTML Test</p>"))
}

func (lc *LotoController) Loto6od45(c echo.Context) error {
	return c.Render(http.StatusOK, "loto/6od45", template.HTML("<p>HTML Test</p>"))
}

func (lc *LotoController) Loto7od39(c echo.Context) error {
	return c.Render(http.StatusOK, "loto/7od39", template.HTML("<p>HTML Test</p>"))
}
