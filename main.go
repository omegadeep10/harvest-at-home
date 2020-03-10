package main

import (
	"log"
	"net/http"

	"harvest-at-home/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()
	models.InitDB("./main.db")

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	r.Get("/items", getItems)
	r.Post("/items", createItem)
	log.Fatal(http.ListenAndServe(":8081", r))
}

type ErrResponse struct {
	Err            error  `json:"-"`               // low-level runtime error
	HTTPStatusCode int    `json:"-"`               // http response status code
	StatusText     string `json:"status"`          // user-level status message
	AppCode        int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText      string `json:"error,omitempty"` // application-level error message, for debugging`
}

type SuccessResponse struct {
	HTTPStatusCode int    `json:"-"`      // http response status code
	StatusText     string `json:"status"` // user-level status message
}

func getItems(w http.ResponseWriter, r *http.Request) {
	items, err := models.AllItems()
	if err != nil {
		render.JSON(w, r, &ErrResponse{
			Err:            err,
			HTTPStatusCode: 500,
			StatusText:     "Something went wrong",
			ErrorText:      err.Error(),
		})
		return
	}

	render.JSON(w, r, items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var i models.Item

	err := render.DecodeJSON(r.Body, i)
	if err != nil {
		render.JSON(w, r, &ErrResponse{
			Err:            err,
			HTTPStatusCode: 400,
			StatusText:     "Unable to parse request body",
			ErrorText:      err.Error(),
		})
		return
	}

	err2 := models.InsertItem(&i)
	if err2 != nil {
		render.JSON(w, r, &ErrResponse{
			Err:            err,
			HTTPStatusCode: 500,
			StatusText:     "Unable to insert item",
			ErrorText:      err.Error(),
		})
		return
	}

	render.JSON(w, r, &SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Successfully created item",
	})
}
