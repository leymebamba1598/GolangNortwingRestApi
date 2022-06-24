package product

//Acceso a la BD, transacciones y consultas

import "database/sql"

type Repository interface{
	GetProductById(productId int)(*Product, error)
	getProducts(params *getProductsRequest)([]*Product, error)
	getTotalProducts()(int,error)
	InsertProduct(params *getAddProductsRequest)(int64, error)
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


	func (repo *repository) getProducts(params *getProductsRequest)([]*Product, error){
	    const sql = `SELECT id, product_code,product_name, coalesce(description,''),standard_cost,list_price,category
		from products order by id limit ? offset ?`

		results, err :=repo.db.Query(sql,params.Limit,params.offset)
		if err != nil {
			panic(err)
		}
		var products []*Product

		for results.Next() { //Recorre tantas veces como registros devuelva el query
			product:=&Product{}
			err=results.Scan(&product.Id, &product.ProductCode, &product.ProductName,&product.Description, 
				&product.StandardCost, &product.ListPrice, &product.Category)
			
			if err != nil {
				panic(err)
			}	
			products=append(products, product)
		}
		return products,nil
	}

	func (repo *repository)	getTotalProducts()(int,error){
		const sql = "select count(*) from products"
		var total int
		row:=repo.db.QueryRow(sql)
		err:=row.Scan(&total)
		if err != nil {
			panic(err)
		}
		return total, nil
	}

	func (repo *repository)	InsertProduct(params *getAddProductsRequest)(int64, error){
		const sql =`insert into products 
		(product_code, product_name,category,description, list_price,standard_cost)
		values(?,?,?,?,?,?)`

		result, err := repo.db.Exec(sql,params.ProductCode,params.ProductName,
			params.Category,params.Description,params.ListPrice,params.StandardCost)
			if(err != nil){
				panic(err)
			}
		
		id,_:= result.LastInsertId()
		return id, nil

	}