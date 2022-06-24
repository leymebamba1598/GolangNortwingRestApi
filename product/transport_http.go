package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

//Llamada a los metodos dependiendo de la accion
func MakeHttpHandler(s Service) http.Handler {
	r:=chi.NewRouter()

    getProductByIdHandler := httptransport.NewServer(makeGetProductByIdEndpoint(s),
		getProductByIdRequestDecoder, //Decodifica parametros de la url
		httptransport.EncodeJSONResponse)

	r.Method(http.MethodGet, "/{id}",getProductByIdHandler)

	getProductsHandler:=httptransport.NewServer(makeGetProductsEndPoint(s),
		getProductsRequestDecoder,
		httptransport.EncodeJSONResponse) //codifica el response
	
	r.Method(http.MethodPost,"/paginated",getProductsHandler)
	
    addProductHandler:= httptransport.NewServer(makeAddProductEndpoint(s),
		addProductRequestDecoder,
		httptransport.EncodeJSONResponse)
	
	r.Method(http.MethodPost, "/",addProductHandler)
	
	return r
}

//Decodifica los parametros del la url
func getProductByIdRequestDecoder(context context.Context, r *http.Request)(interface{}, error){
	ProductId,_:=strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIdRequest{
		ProductId:ProductId,
	},nil
}

//Decodifica los parametros del cuerpo body
func getProductsRequestDecoder(context context.Context, r *http.Request)(interface{}, error){
	request :=getProductsRequest{}
	err:=json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request,nil
}

func addProductRequestDecoder(context context.Context, r *http.Request)(interface{}, error){
   request:= getAddProductsRequest{}
   err:=json.NewDecoder(r.Body).Decode(&request)
   if err != nil {
	panic(err)
  }
  return request,nil
}