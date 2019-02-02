package server

import "net/http"
import "github.com/bakhtik/webapp/internal/app/handler"

func (s *Server) routes() {
	handler := handler.NewHandler(s.db, s.cache)
	files := http.FileServer(http.Dir("web/public"))
	s.mux.Handle("/static/", http.StripPrefix("/static/", files))
	s.mux.Handle("/favicon.ico", http.NotFoundHandler())
	s.mux.Handle("/", handler.Index())
}
