package resource

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler(
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recover)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Route("/",func(api chi.Router){
		api.Post("transactions")
	})
	return r 
)
