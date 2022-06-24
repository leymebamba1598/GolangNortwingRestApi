package product

//Logica de negocio

type Service interface{
	GetProductById(param *getProductByIdRequest)(*Product, error)
	GetProducts(param *getProductsRequest)(*ProductList,error)
}

type service struct {
	repo Repository
}
func NewService(repo Repository)Service {
	return &service{
		repo: repo,
	}
}
func (s *service)GetProductById(param *getProductByIdRequest)(*Product, error){
	//Logica
		/*products, err:=s.repo.GetProductById(param.ProductId)
		if err != nil{
			panic(err)
		}
		return &Product{Id: products.Id,ProductCode: products.ProductCode,ProductName :products.ProductName,  Description:products.Description, StandardCost:products.StandardCost, ListPrice :products.ListPrice,   Category:products.Category}, nil
		*/
		return s.repo.GetProductById(param.ProductId)
}
func(s *service) GetProducts(params *getProductsRequest)(*ProductList,error){
		products, err:= s.repo.getProducts(params)
		if err != nil{
			panic(err)
		}
		totalProducts, err := s.repo.getTotalProducts()
		if err != nil{
			panic(err)
		}
		return  &ProductList{Data: products,TotalRecords: totalProducts},nil
		
}