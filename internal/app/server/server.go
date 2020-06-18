package server

import (
	"net/http"
	"os"

	"github.com/beldmian/ORS/config"
	"github.com/beldmian/ORS/internal/app/db"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	db     db.Datebase
	config config.SeverConfig
	router *mux.Router
	logger *logrus.Logger
}

// New ...
func New(config config.SeverConfig, db db.Datebase) Server {
	return Server{
		db:     db,
		config: config,
		router: mux.NewRouter(),
		logger: logrus.New(),
	}
}

// Start ...
func (s Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("Starting server...")

	s.configureRouter()
	s.logger.Debug("router started successful on port: ", s.config.Port)
	return http.ListenAndServe(s.config.Port, s.router)
}

func (s Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.Out = os.Stderr
	switch s.config.LogType {
	case "stdout":
		s.logger.Out = os.Stdout
	case "file":
		file, err := os.OpenFile(s.config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			s.logger.Warn("Log file not works: ", err)
		} else {
			s.logger.Out = file
		}
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *Server) configureRouter() {
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(
		http.Dir("./internal/app/server/static/"))))
	s.router.HandleFunc("/", s.indexHandler)
}
