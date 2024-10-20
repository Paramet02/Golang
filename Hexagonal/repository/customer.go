package repository

// Repository manage database
// การกำหนดข้อมูล ภายใน struct 
type Customer struct {
	CustomerID  int    `gorm:"column:customer_id"`
	Name        string `gorm:"column:name"`
	DateOfBirth string `gorm:"column:date_of_birth"`
	City        string `gorm:"column:city"`
	ZipCode     string `gorm:"column:zipcode"`
	Status      int    `gorm:"column:status"`
}

// port = interface 
// กำหนดสเปกของ port
type CustomerRepository interface {
	GetAll() ([]Customer , error)
	GetById(int) (*Customer , error)
} 