package main

import (
	"fmt"
	"net/http"
	"github.com/hannanmiah/golang-tutorial/handlers"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	http.HandleFunc("GET /products", handlers.GetProducts)
	http.HandleFunc("POST /products", handlers.AddProduct)
	http.HandleFunc("GET /products/{id}", handlers.GetProduct)
	http.HandleFunc("PUT /products/{id}", handlers.UpdateProduct)
	http.HandleFunc("DELETE /products/{id}", handlers.DeleteProduct)

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}