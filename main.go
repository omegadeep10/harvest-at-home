package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

var tokenAuth *jwtauth.JWTAuth = jwtauth.New("HS256", []byte("secret"), nil)

func main() {
	r := chi.NewRouter()

	_, tokenStr, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenStr)

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		jwtauth.Verifier(tokenAuth),
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Authenticator)

		r.Get("/hello", getHello)
	})

	log.Fatal(http.ListenAndServe(":8081", r))
}

type Hello struct {
	Title string `json:"title"`
}

func getHello(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	fmt.Printf("protected area. hi %v", claims["user_id"])

	h := Hello{Title: "World"}

	render.JSON(w, r, h)
}
