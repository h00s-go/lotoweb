package templates

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// Template represect Template struct
type Template struct {
	templates *template.Template
}

// LoadTemplates load all templates
func LoadTemplates() *Template {
	return &Template{templates: template.Must(template.ParseFiles(
		"templates/layout/header.html",
		"templates/layout/navigation.html",
		"templates/layout/footer.html",

		"templates/loto/index.html",
	))}
}

// Render for templates
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
