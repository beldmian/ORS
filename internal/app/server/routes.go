package server

import (
	"html/template"
	"net/http"

	"github.com/beldmian/ORS/internal/app/types"
	"github.com/gorilla/mux"
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
	if err := tmpl.ExecuteTemplate(w, "events", struct {
		Title  string
		Events []types.Event
	}{
		Title:  "ORS",
		Events: events,
	}); err != nil {
		s.logger.Fatal(err)
	}
}

func (s Server) eventHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	tmpl := template.Must(template.ParseFiles(
		"./internal/app/server/templates/event.html",
		"./internal/app/server/templates/layouts/header.html",
		"./internal/app/server/templates/layouts/footer.html"))
	event, err := s.db.GetEvent(name)
	if err != nil {
		s.logger.Fatal("Cannot get event: ", err)
	}
	if err := tmpl.ExecuteTemplate(w, "event", struct {
		Title string
		Event types.Event
	}{
		Title: name,
		Event: event,
	}); err != nil {
		s.logger.Fatal(err)
	}
}
