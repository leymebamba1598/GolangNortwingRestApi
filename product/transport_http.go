package product

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

//Llamada a los metodos dependiendo de la accion
func MakeHttpHandler(s Service) http.Handler {
	r:=chi.NewRouter()

    getProductByIdHandler := httptransport.NewServer(makeGetProductByIdEndpoint(s),
		getProductByIdRequestDecoder, 
		httptransport.EncodeJSONResponse)

	r.Method(http.MethodGet, "/{id}",getProductByIdHandler)
	return r
}

func getProductByIdRequestDecoder(context context.Context, r *http.Request)(interface{}, error){
	ProductId,_:=strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIdRequest{
		ProductId:ProductId,
	},nil
}