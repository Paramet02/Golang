package service

// business database
// data transfer object DTO
// ปั้นข้อมูลเพื่อ ฝั้งผู้ใช้งาน ไม่รับรู้อะไรมากกว่านี้
type CustomerResponse struct {
	CustomerID  int    `json:"customer_id"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
}

// port 
type CustomerService interface {
	GetCustomers() ([]CustomerResponse , error)
	GetCustomer(int) (*CustomerResponse , error)
}
