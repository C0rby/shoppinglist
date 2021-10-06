package api

import (
	"net/http"
	"strings"

	"github.com/c0rby/shoppinglist/pkg/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Handler(service service.Service) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/v1/lists", func(w http.ResponseWriter, r *http.Request) {
		lists, _ := service.GetLists()
		var sb strings.Builder
		for _, l := range lists {
			sb.WriteString(l.Name)
			sb.WriteRune('\n')
		}
		w.Write([]byte(sb.String()))
	})

	return r
}
