package server

import (
	"net/http"

	"github.com/c0rby/shoppinglist/pkg/api"
	"github.com/c0rby/shoppinglist/pkg/service"
	"github.com/c0rby/shoppinglist/pkg/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	address string
}

func New() Server {
	return Server{address: ":8000"}
}

func (s Server) ListenAndServe() error {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodOptions, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders: []string{"Content-Type"},
	}))
	db, err := storage.NewSqlite3DB("./shoppinglist.sqlite?_fk=true")
	if err != nil {
		return err
	}
	defer db.Close()
	if err := storage.CreateTables(db); err != nil {
		return err
	}
	store, err := storage.NewSqlStore(db)
	if err != nil {
		return err
	}
	service := service.New(store)
	router.Mount("/api", api.Handler(service))

	return http.ListenAndServe(s.address, router)
}
