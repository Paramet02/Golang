package service

// Product struct
// โครงสร้างสำหรับเก็บข้อมูลสินค้า
type Product struct {
	ID       int    `json:"id"`// รหัสสินค้า
	Name     string `json:"name"`// ชื่อสินค้า
	Quantity int    `json:"quantity"`// จำนวนสินค้า
}

// CatalogService interface
// ประกาศเมธอด GetProduct เพื่อใช้ดึงข้อมูลสินค้า
// ทำให้ struct ไหนก็ตามที่ใช้ interface นี้ต้องมีฟังก์ชัน GetProduct
type CatalogService interface {
	GetProduct() ([]Product, error) // ฟังก์ชันสำหรับดึงข้อมูลสินค้า
}