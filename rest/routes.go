package rest

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
		mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
	))

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProduct),
	))

	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(handlers.GetProduct),
	))

	mux.Handle("DELETE /products/{id}",manager.With(
		http.HandlerFunc(handlers.DeleteProduct),
	))
	mux.Handle("PUT /products/{id}",manager.With(
		http.HandlerFunc(handlers.UpdateProduct),
	))

}