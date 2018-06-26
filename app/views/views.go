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
func LoadViews() *View {
	return &View{views: template.Must(template.ParseFiles(
		"app/views/layout/header.html",
		"app/views/layout/navigation.html",
		"app/views/layout/footer.html",

		"app/views/loto/index.html",
	))}
}

// Render for templates
func (v *View) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return v.views.ExecuteTemplate(w, name, data)
}
