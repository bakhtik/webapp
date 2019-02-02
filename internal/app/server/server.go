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
	config     *Config
}

type Config struct {
	DbConnString string
	CacheAddr    string
	ServerAddr   string
}

func NewServer(config *Config) *Server {
	s := &Server{
		// just in case you need some setup here
		config: config,
	}
	return s
}

func (s *Server) Start() {
	// Startup all dependencies
	s.Dependencies()

	// Startup the http Server in a way that
	// we can gracefully shut it down again
	s.mux = http.NewServeMux()
	// register handlers
	s.routes()

	s.httpServer = &http.Server{Addr: s.config.ServerAddr, Handler: s.mux}
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

func (s *Server) Dependencies() {
	db, err := models.NewDB(s.config.DbConnString)
	if err != nil {
		log.Panic(err)
	}
	s.db = db
	s.cache = models.NewClient(s.config.CacheAddr)
}
