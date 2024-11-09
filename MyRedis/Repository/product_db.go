package repository

import (
	
	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewproductRepository(db *gorm.DB) ProductRepo {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoryDB{db}
}



func (p productRepositoryDB) CreateProduct(product product) error {
    err := p.db.Create(&product).Error // เก็บข้อผิดพลาดไว้ในตัวแปร err
    if err != nil {
        return err // ถ้ามีข้อผิดพลาดก็คืนค่าข้อผิดพลาดนั้น
    }
    return nil // ถ้าไม่มีข้อผิดพลาดก็คืนค่า nil
}

func (p productRepositoryDB) UpdateProduct(product product) error {
	// ใช้ Updates แทน Update เพื่ออัปเดตหลายๆ fields พร้อมกัน
	err := p.db.Model(&product).Where("id = ?", product.ID).Updates(map[string]interface{}{
		"name":     product.Name,
		"quantity": product.Quantity,
	}).Error
	
	if err != nil {
		return err
	}

	return nil
}

func (p productRepositoryDB) GetProducts() (pro []product , err error) {
	// select * from products order by quantity desc limit 15 
	err = p.db.Order("quantity desc").Limit(15).Find(&pro).Error
	if err != nil {
        return nil , err // ถ้ามีข้อผิดพลาดก็คืนค่าข้อผิดพลาดนั้น
    }
	return pro , nil
}

func (p productRepositoryDB) GetProduct(name string) (pro []product , err error) {
	// select * product Where name = %name%
	err = p.db.Where("name LIKE ?", "%" + name + "%").First(&pro).Error
	if err != nil {
		return nil , err
	}

	return pro , nil
}