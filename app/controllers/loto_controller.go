package controllers

import (
	"html/template"
	"net/http"

	"github.com/h00s/lotoweb/app/models"
	"github.com/labstack/echo"
)

// LotoController for Loto methods
type LotoController struct {
	lotoModel *models.LotoModel
}

// NewLotoController creates new loto controller
func NewLotoController() *LotoController {
	return &LotoController{models.NewLotoModel()}
}

// Index is first page
func (lc *LotoController) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "loto/index", template.HTML(""))
}

// Loto6od45 return numbers for loto 6 od 45
func (lc *LotoController) Loto6od45(c echo.Context) error {
	return c.Render(http.StatusOK, "loto/6od45", lc.lotoModel.GetNumbers(6, 45))
}

// Loto7od39 return numbers for loto 7 od 39
func (lc *LotoController) Loto7od39(c echo.Context) error {
	return c.Render(http.StatusOK, "loto/7od39", lc.lotoModel.GetNumbers(7, 39))
}
