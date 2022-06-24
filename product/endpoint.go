package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//Entidades que van almacenar parametros de la url o en el body, todo lo que entra

type getProductByIdRequest struct {
	ProductId int 
}
type getProductsRequest struct {
	Limit int
	offset int
}

type getAddProductsRequest struct {
	Category string
	Description string
	ListPrice string
	StandardCost string
	ProductCode string
	ProductName string
}

//Convierte el request y llama al servicio
func makeGetProductByIdEndpoint(s Service) endpoint.Endpoint{
	getProductByIdEndpoint:=func(ctx context.Context, request interface{})(interface{},error){
		req:=request.(getProductByIdRequest)
         product, err :=s.GetProductById(&req)
		 if err != nil {
			panic(err)
		 }
		 return product, nil
	} 
 return getProductByIdEndpoint
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint{
	getProductsEndPoint:=func(ctx context.Context, request interface{})(interface{},error){
	    req:=request.(getProductsRequest) //convertimos el request al tipo getProductsRequest
		result, err :=s.GetProducts(&req)
		if err != nil {
			panic(err)
		}
		return result,nil
	}

	return getProductsEndPoint
}

func makeAddProductEndpoint(s Service) endpoint.Endpoint{
		addproductEndpoint:=func(ctx context.Context, request interface{})(interface{},error){
			req:=request.(getAddProductsRequest)
			productId, err :=s.InsertProduct(&req)
			if err != nil {
				panic(err)
			}
			return productId, nil
		}
		return addproductEndpoint
}