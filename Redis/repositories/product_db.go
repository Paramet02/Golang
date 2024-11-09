package repositories

import (
	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

// ProductRepository คือตัวกำหนดให้ adapter ตัวให้คอนเฟิร์ฒตาม ProductRepository
func NewproductRepositoryDB(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoryDB{db : db}
}

func (p productRepositoryDB) GetProduct() (pro []product , err  error) {
	err = p.db.Order("quantity desc").Limit(15).Find(&pro).Error
	return pro , err
}