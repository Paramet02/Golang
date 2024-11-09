package repository

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type product struct {
	ID int
	Name string
	Quantity int
}

type ProductRepo interface{
	CreateProduct(product product) error
	UpdateProduct(product product) error
	GetProducts() ([]product , error) 
	GetProduct(name string) ([]product , error)
} 


func mockData(db *gorm.DB) error {

	var count int64 
	// all product count
	db.Model(&product{}).Count(&count)
	if count > 0 {
		return nil
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	products := []product{}
	for i := 0 ; i < 5000 ; i++ {
		products = append(products, product{
			Name: fmt.Sprintf("Product %v " , i+1),
			Quantity: random.Intn(100),
		})
	}
	return db.Create(&products).Error
}

// เพิ่มฟังก์ชันนี้ใน repository package
func NewProduct(name string, quantity int) product {
	return product{
		Name:     name,
		Quantity: quantity,
	}
}

func NewUpdateProduct(id int, name string, quantity int) product {
    return product{
        ID:       id,
        Name:     name,
        Quantity: quantity,
    }
}