package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	r.Get("/", getHello)
	log.Fatal(http.ListenAndServe(":8081", r))
}

type Hello struct {
	Title string `json:"title"`
}

func getHello(w http.ResponseWriter, r *http.Request) {
	h := Hello{Title: "World"}

	render.JSON(w, r, h)
}
