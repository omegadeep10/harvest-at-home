package controllers

import (
	"encoding/json"
	"harvest-at-home/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := models.GetAllItems()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &models.ErrResponse{
			HTTPStatusCode: 500,
			ErrorText:      err.Error(),
		})
		return
	}

	render.JSON(w, r, items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &models.ErrResponse{
			HTTPStatusCode: 400,
			ErrorText:      err.Error(),
		})
		return
	}

	err2 := models.CreateItem(&item)
	if err2 != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &models.ErrResponse{
			HTTPStatusCode: 500,
			ErrorText:      err2.Error(),
		})
		return
	}

	render.JSON(w, r, &models.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Successfully created item",
	})
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	item_id := chi.URLParam(r, "id")
	err := models.DeleteItem(item_id)

	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, &models.ErrResponse{
			HTTPStatusCode: 404,
			ErrorText:      err.Error(),
		})
		return
	}

	render.JSON(w, r, &models.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Successfully deleted item " + item_id,
	})
}
