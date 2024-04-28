package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

func Load_Templates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func Exec_Template(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
