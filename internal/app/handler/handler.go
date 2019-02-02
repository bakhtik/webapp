package handler

import "github.com/bakhtik/webapp/internal/app/models"

type Handler struct {
	db    models.Datastore
	cache models.Cache
}

func NewHandler(db models.Datastore, cache models.Cache) *Handler {
	return &Handler{
		db:    db,
		cache: cache,
	}
}
