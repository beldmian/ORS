package server

import (
	"html/template"
	"net/http"
)

func (s Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"./internal/app/server/templates/index.html",
		"./internal/app/server/templates/layouts/header.html",
		"./internal/app/server/templates/layouts/footer.html"))
	if err := tmpl.ExecuteTemplate(w, "index", map[string]string{
		"Title": "ORS",
	}); err != nil {
		s.logger.Fatal(err)
	}
}
