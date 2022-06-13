package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//Entidades que van almacenar parametros de la url o en el body

type getProductByIdRequest struct {
	ProductId int 
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