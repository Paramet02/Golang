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

// คือการกำหนด adapter แต่ล่ะตัวต้องมีไรบ้าง
type ProductRepo interface{
	CreateProduct(product product) error
	UpdateProduct(product product) error
	GetProducts() ([]product , error) 
	GetProduct(name string) ([]product , error)
} 

// คือการจำลอง data 
func mockData(db *gorm.DB) error {
	
	// ตรวจสอบจำนวนสินค้าทั้งหมดในฐานข้อมูลก่อน หากมีข้อมูลแล้วจะไม่เพิ่มข้อมูลใหม่
	var count int64 
	// all product count
	// model ไว้ใช้กับ struct 
	db.Model(&product{}).Count(&count)
	if count > 0 {
		return nil
	}

	// สร้างข้อมูลจำลอง (mock data) เพื่อทดสอบ
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	// เพิ่มข้อมูลสินค้า 5000 รายการ
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
	// ฟังก์ชันนี้ใช้ในการสร้างข้อมูลสินค้าใหม่ โดยรับค่า name และ quantity
	// ส่งคืนเป็น struct product ที่พร้อมใช้ใน repository

	return product{
		Name:     name,
		Quantity: quantity,
	}
}

func NewUpdateProduct(id int, name string, quantity int) product {
	// ฟังก์ชันนี้ใช้ในการสร้างข้อมูลสินค้าใหม่ที่มีการอัปเดต
	// โดยรับค่า id, name และ quantity เพื่อใช้ในการอัปเดตข้อมูลในฐานข้อมูล
	
    return product{
        ID:       id,
        Name:     name,
        Quantity: quantity,
    }
}