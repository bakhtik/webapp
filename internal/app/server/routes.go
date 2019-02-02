package server

import "net/http"

func (s *Server) routes() {
	files := http.FileServer(http.Dir("web/public"))
	s.mux.Handle("/static/", http.StripPrefix("/static/", files))
	s.mux.Handle("/favicon.ico", http.NotFoundHandler())
}
