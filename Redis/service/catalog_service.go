package service

import "paramet/repositories"

type catalogService struct {
	productRepo repositories.ProductRepository
}

func NewcatalogService( productRepo repositories.ProductRepository ) CatalogService {
	return catalogService{productRepo}
}

func (c catalogService) GetProduct() (pro []Product, err error) {
	productsdb , err := c.productRepo.GetProduct()
	if err != nil {
		return nil, err
	}

	for _ , p := range productsdb {
		pro = append(pro, Product{
			ID: p.ID,
			Name: p.Name,
			Quantity: p.Quantity,
		})
	}
	return pro , nil
}