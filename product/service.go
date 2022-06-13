package product

//Logica de negocio

type Service interface{
	GetProductById(param *getProductByIdRequest)(*Product, error)
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
		return s.repo.GetProductById(param.ProductId)
}