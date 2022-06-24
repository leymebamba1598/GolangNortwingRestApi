package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leymebamba1598/GolangNortwingRestApi/database"
	"github.com/leymebamba1598/GolangNortwingRestApi/product"
)

func main() {

	databaseConnection:=database.InitDB()
	defer databaseConnection.Close() 
	
	var productRepository=product.NewRepository(databaseConnection)
	
	var productService product.Service
	productService=product.NewService(productRepository)
	
	r:=chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))
	http.ListenAndServe(":3000", r)


}


