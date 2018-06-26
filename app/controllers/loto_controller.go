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
	return c.Render(http.StatusOK, "index", template.HTML("<p>HTML Test</p>"))
}
