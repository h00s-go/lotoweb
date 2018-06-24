package loto

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo"
)

// Controller for Link methods
type Controller struct {
}

// NewController creates new link controller
func NewController() *Controller {
	return &Controller{}
}

// Index is first page
func (lc *Controller) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", template.HTML("<p>HTML Test</p>"))
}
