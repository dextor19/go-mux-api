package product

import (
	"github.com/gorilla/mux"

)
func RegisterRoutes(router *mux.Router) {
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/products", GetProducts).Methods("GET")
	v1.HandleFunc("/products/{id}", GetProductById).Methods("GET")
	v1.HandleFunc("/products", CreateProduct).Methods("POST")
	v1.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	v1.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
}
