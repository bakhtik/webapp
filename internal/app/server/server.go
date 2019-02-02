package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bakhtik/webapp/internal/app/models"
)

type Server struct {
	db         models.Datastore
	cache      models.Cache
	mux        *http.ServeMux
	httpServer *http.Server
}

type Config struct {
	dbConnString string
	cacheAddr    string
	serverAddr   string
}

func NewServer() *Server {
	s := &Server{
		// just in case you need some setup here
	}
	return s
}

func (s *Server) Start(config Config) {
	// Startup all dependencies
	s.Dependencies(config)

	// Startup the http Server in a way that
	// we can gracefully shut it down again
	s.mux = http.NewServeMux()
	// register handlers
	s.routes()

	s.httpServer = &http.Server{Addr: config.serverAddr, Handler: s.mux}
	err := s.httpServer.ListenAndServe() // Blocks!
	if err != http.ErrServerClosed {
		log.Print("Http Server stopped unexpected")
		s.Shutdown()
	} else {
		log.Print("Http Server stopped")
	}
}

func (s *Server) Shutdown() {
	if s.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			log.Print("Failed to shutdown http server gracefully")
		} else {
			s.httpServer = nil
		}
	}
}

func (s *Server) Dependencies(config Config) {
	db, err := models.NewDB(config.dbConnString)
	if err != nil {
		log.Panic(err)
	}
	s.db = db
	s.cache = models.NewClient(config.cacheAddr)
}
