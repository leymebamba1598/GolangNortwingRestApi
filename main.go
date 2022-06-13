package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leymebamba1598/GolangNortwingRestApi/database"
)

var databaseConnection *sql.DB

type Product struct {
	ID           int    `json:"codigo"`
	Product_Code string `json:"codigo_producto"`
	Description  string `json:"descripcion"`
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {

	databaseConnection = database.InitDB()
	defer databaseConnection.Close() //se ejecuta cuando finaliza la funcion

	r := chi.NewRouter()
	//Rutas
	r.Get("/products", AllProduct)
	r.Post("/products", CreateProducts)
	r.Put("/products/{id}", UpdateProducts)
	r.Delete("/products/{id}", DeleteProduct)
	http.ListenAndServe(":3000", r)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    id:=chi.URLParam(r,"id")

	query, err := databaseConnection.Prepare("delete from products where id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	defer query.Close()
	responseWhitJson(w, http.StatusCreated, map[string]string{"message": "Delete correctamente"})
}

func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	var producto Product
    id:=chi.URLParam(r,"id")
	json.NewDecoder(r.Body).Decode(&producto)

	query, err := databaseConnection.Prepare("update products set product_code=?,description=? where id=?")
	catch(err)
	_, er := query.Exec(producto.Product_Code, producto.Description,id)
	catch(er)
	defer query.Close()
	responseWhitJson(w, http.StatusCreated, map[string]string{"message": "Update correctamente"})
}

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var producto Product
	json.NewDecoder(r.Body).Decode(&producto)
	query, err := databaseConnection.Prepare("Insert products SET product_code=?, description=?")
	catch(err)
	_, er := query.Exec(producto.Product_Code, producto.Description)
	catch(er)
	defer query.Close()
	responseWhitJson(w, http.StatusCreated, map[string]string{"message": "Creador correctamente"})
}

func AllProduct(w http.ResponseWriter, r *http.Request) {
	const sql = `select id, product_code,coalesce(description,'') 
	as description from products`
	results, err := databaseConnection.Query(sql)
	catch(err)
    
	var lstProductos []*Product //Puntero de tipo slice de products

	for results.Next() {
		producto := &Product{}
		err := results.Scan(&producto.ID, &producto.Product_Code, &producto.Description)
		//Scan, scanea toda las propiedades resultantes del script, y lo mapeamos
		if err != nil {
			panic(err)
		}

		lstProductos = append(lstProductos, producto)
	}
	responseWhitJson(w, http.StatusOK, lstProductos)

}

func responseWhitJson(w http.ResponseWriter, code int, payload interface{}) {
	//payload tipo interace vacia, se convierte al tipo de dato que se lo envie map, slice, entero,etc
	response, _ := json.Marshal(payload)               // se convierte a json
	w.Header().Set("Content-Type", "application/json") //respuesta estara en formato json
	w.WriteHeader(code)
	w.Write(response)
}
