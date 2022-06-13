package product

//Acceso a la BD, transacciones y consultas

import "database/sql"

type Repository interface{
	GetProductById(productId int)(*Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB)Repository{ //Creamos la conexion
	return &repository{db: databaseConnection}
}

func (repo *repository) GetProductById(productId int)(*Product, error){
	const sql = `SELECT id, product_code,product_name, coalesce(description,''),standard_cost,list_price,category
					from products where id=?`
    
	row:=repo.db.QueryRow(sql,productId)
	product:= &Product{}

    err := row.Scan(&product.Id, &product.ProductCode,&product.ProductName,
		&product.Description,&product.StandardCost,&product.ListPrice,&product.Category) //Mapeamos los resultados de la consulta en el producto
   
      if err != nil {
            panic(err)
	  }
	  return product, err
	}