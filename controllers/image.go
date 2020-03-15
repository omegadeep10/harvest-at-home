package controllers

import (
	"fmt"
	"harvest-at-home/models"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/go-chi/render"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("image")
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &models.ErrResponse{
			HTTPStatusCode: 500,
			ErrorText:      err.Error(),
		})
		return
	}

	defer file.Close()
	fmt.Println("Uploaded file: ", handler.Filename, " - File Size: ", handler.Size, " - MIME Header: ", handler.Header)

	tempFile, err2 := ioutil.TempFile("static", "upload-*"+filepath.Ext(handler.Filename))
	if err2 != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &models.ErrResponse{
			HTTPStatusCode: 500,
			ErrorText:      err2.Error(),
		})
		return
	}

	defer tempFile.Close()

	fileBytes, err3 := ioutil.ReadAll(file)
	if err3 != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &models.ErrResponse{
			HTTPStatusCode: 500,
			ErrorText:      err3.Error(),
		})
		return
	}

	tempFile.Write(fileBytes)
	finfo, _ := tempFile.Stat()

	render.JSON(w, r, &models.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     finfo.Name(),
	})
}
