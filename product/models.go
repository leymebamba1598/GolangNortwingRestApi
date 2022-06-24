package product

//Mapear los valores que nos devuelve la BD, Todo lo que sale
type Product struct {
	Id int `json:"id"`
	ProductCode string  `json:"productCode"`
	ProductName string   `json:"productName"`
	Description string  `json:"description"`
	StandardCost float64   `json:"standartCost"` 
	ListPrice float64  `json:"listPrice"` 
	Category string  `json:"category"`
}

type ProductList struct {
	Data[]*Product `json:"data"`
	TotalRecords int `json:"totalRecords"`
}