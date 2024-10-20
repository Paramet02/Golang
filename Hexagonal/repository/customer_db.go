// The code defines a customer repository struct for interacting with a database using GORM in Go.
// repository ส่งออกให้ service จัดการ และทำให้ handler code clean ที่สุด
package repository

import (
	"gorm.io/gorm"
)

// adapter = struct
type customerRepositoryDB struct {
	db *gorm.DB // Do not allow direct access
}

// new instance
// port + adapter
func NewCustomerRepositoryDB(db *gorm.DB) CustomerRepository {
	return &customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	
	// SELECT * FROM customers
	if result := r.db.Find(&customers); result.Error != nil {
		// Handle database errors
		return nil, result.Error
	}
	
	return customers, nil
}

func (r *customerRepositoryDB) GetById(id int) (*Customer, error) {
	customers := Customer{}

	// SELECT * FROM customers WHERE id = ?
	if result := r.db.First(&customers , id); result.Error != nil {
		// Handle database errors
		return nil , result.Error
	}

	return &customers, nil
}
