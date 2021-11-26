package handlers

import (
	"html/template"
)

func parseHtmlTemplate(templatePath string) *template.Template {
	tmpl := template.Must(template.ParseFiles(templatePath))
	return tmpl
}
