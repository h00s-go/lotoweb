package views

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// View represent views struct
type View struct {
	views *template.Template
}

// LoadViews load all templates
func LoadViews(files []string) *View {
	return &View{views: template.Must(template.ParseFiles(files...))}
}

// Render for templates
func (v *View) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return v.views.ExecuteTemplate(w, name, data)
}
