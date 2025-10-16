package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type product struct {
	Name string
	ID   int
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		productName := r.URL.Query().Get("productName")
		productColor := r.URL.Query().Get("productColor")

		if productName != "" {
			w.Write([]byte(productName))
		} else if productColor != "" {
			w.Write([]byte(productColor))
		} else {
			w.Write([]byte("Products List"))
		}
	})

	r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product product

		err := render.DecodeJSON(r.Body, &product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product.ID = 5

		render.JSON(w, r, product)
	})

	http.ListenAndServe(":3000", r)
}

// func myMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		println("Before")
// 		next.ServeHTTP(w, r)
// 		println("After")
// 	})
// }

// func myMiddleware2(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		println("request method ", r.Method, " resource ", r.URL.Path)
// 		next.ServeHTTP(w, r)
// 		println("After 2")
// 	})
// }
