package repository

import "errors"

// ในที่นี้ customerRepositoryMock ถูกใช้ใน CustomerService เพื่อจำลองข้อมูลลูกค้า 
// ทำให้คุณสามารถทดสอบการดึงข้อมูลลูกค้าหรือแสดงข้อมูลได้โดยไม่ต้องเชื่อมต่อกับฐานข้อมูลจริง.
type customerRepositoryMock struct {
	customers []Customer
}

// ถูกสร้างมาเป็น data จำลอง เลยไม่ต้องรับค่า
func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{CustomerID: 1001, Name: "Ashish", City: "New Delhi", ZipCode: "110011", DateOfBirth: "2000-01-01", Status: 1},
		{CustomerID: 1002, Name: "Rob", City: "New Delhi", ZipCode: "110011", DateOfBirth: "2000-01-01", Status: 0},
	}

	// {customers: customers} คือการใส่ข้อมูล ภายในข้างนอกไปยังภายใน customerRepositoryMock
	return &customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _ , customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}


	return nil, errors.New("customer not fonud")
}