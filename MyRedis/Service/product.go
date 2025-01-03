package service


type Product struct {
	ID       int    `json:"id"`// รหัสสินค้า
	Name     string `json:"name"`// ชื่อสินค้า
	Quantity int    `json:"quantity"`// จำนวนสินค้า
}

// คือการกำหนด adapter แต่ล่ะตัวต้องมีไรบ้าง
type ProductService interface{
	CreateProduct(product Product) error
	UpdateProduct(product Product) error
	GetProducts() ([]Product , error)
	GetProduct(name string) ([]Product , error)
} // c