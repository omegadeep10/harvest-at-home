package main

import (
	"log"
	"net/http"

	"harvest-at-home/controllers"
	"harvest-at-home/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()
	models.InitDB("./main.db")
	defer models.CloseDB()

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	r.Get("/items", controllers.GetItems)
	r.Post("/items", controllers.CreateItem)
	r.Delete("/items/{id}", controllers.DeleteItem)

	log.Fatal(http.ListenAndServe(":8081", r))
}
