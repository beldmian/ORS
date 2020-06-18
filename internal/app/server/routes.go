package server

import (
	"html/template"
	"net/http"

	"github.com/beldmian/ORS/internal/app/types"
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

func (s Server) eventsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"./internal/app/server/templates/events.html",
		"./internal/app/server/templates/layouts/header.html",
		"./internal/app/server/templates/layouts/footer.html"))
	events, err := s.db.GetEvents()
	if err != nil {
		s.logger.Fatal("Cannot get events: ", err)
	}
	if err := tmpl.ExecuteTemplate(w, "index", struct {
		Title  string
		Events []types.Event
	}{
		Title:  "ORS",
		Events: events,
	}); err != nil {
		s.logger.Fatal(err)
	}
}
