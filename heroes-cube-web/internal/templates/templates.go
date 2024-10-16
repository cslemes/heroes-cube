package templates

import (
	"html/template"
	"net/http"
)

var (
	homeTemplate       *template.Template
	charactersTemplate *template.Template
)

func Init() error {
	var err error

	homeTemplate, err = template.ParseFiles("internal/templates/layout.html", "internal/templates/home.html")
	if err != nil {
		return err
	}

	charactersTemplate, err = template.ParseFiles("internal/templates/layout.html", "internal/templates/characters.html")
	if err != nil {
		return err
	}

	return nil
}

func RenderHome(w http.ResponseWriter) error {
	return homeTemplate.ExecuteTemplate(w, "layout", nil)
}

func RenderCharacters(w http.ResponseWriter, data interface{}) error {
	return charactersTemplate.ExecuteTemplate(w, "layout", data)
}
