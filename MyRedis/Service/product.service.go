package service

import repository "github.com/paramet02/Repository"

type proService struct {
	proRepo repository.ProductRepo
}

func NewproService(proRepo repository.ProductRepo) ProductService {
	return proService{proRepo}
}

func (s proService) CreateProduct(product Product) error {
	// แปลงจาก service.Product ไปเป็น repository.product โดยใช้ NewProduct
	repoProduct := repository.NewProduct(product.Name, product.Quantity)

	// เรียกใช้ CreateProduct ของ repository
	return s.proRepo.CreateProduct(repoProduct)
}

func (s proService) UpdateProduct(product Product) error {
	repoProduct := repository.NewUpdateProduct(product.ID, product.Name, product.Quantity)
	err := s.proRepo.UpdateProduct(repoProduct)
	if err != nil {
		return err
	}
	return nil
}

func (s proService) GetProducts() (pro []Product, err error) {
	products, err := s.proRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		pro = append(pro, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	return pro, nil
}

func (s proService) GetProduct(name string) (pro []Product, err error) {
	newproduct, err := s.proRepo.GetProduct(name)
	if err != nil {
		return nil, err
	}
	for _, p := range newproduct {
		pro = append(pro, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	return pro, nil
}