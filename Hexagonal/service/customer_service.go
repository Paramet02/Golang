package service

//  business logic
// service คือ คนคุมควมทุกอย่าง เช่นจัดการ error 
import (
	"github.com/gofiber/fiber/v2"
	"paramet/logs"
	"paramet/repository"

	"gorm.io/gorm"
)

// adapter
type customerService struct {
	cr repository.CustomerRepository
}
//  business logic = service port + repository port 
func NewCustomerService(cr repository.CustomerRepository) CustomerService {
	return &customerService{cr : cr}
}

func (a customerService) GetCustomers() ([]CustomerResponse , error) {
	// ฟังก์ชัน GetAll() ใน repository ถูกออกแบบมาเพื่อดึงข้อมูลจากฐานข้อมูลและส่งข้อมูลนั้นกลับไปยังผู้เรียกใช้ เช่น service
	customers , err := a.cr.GetAll()

	if err != nil {
		logs.Error(err)
		// ส่ง error ไปยัง handler เพื่อให้ handler ไม่ต้องรับรู้อะไร
		return nil , fiber.NewError(fiber.StatusInternalServerError , "unexpected error")
	}
	// ข้อมูลเปล่า
	cResponses := []CustomerResponse{}
	// ข้อมูลที่กำลังจะปั้น
	for _ , customer := range customers {
		cResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name: customer.Name,
			Status: customer.Status,
		}
		// cResponses <- cResponse
		cResponses = append(cResponses , cResponse)
	}
	
	return cResponses , nil 
}

func (a customerService) GetCustomer(id int) (*CustomerResponse , error) {
	// return เป็น value กับ err แล้วเอาข้อมูลที่ได้ มาทำการปั้นข้อมูล
	customers , err := a.cr.GetById(id)
	
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			// ส่ง error ไปยัง handler เพื่อให้ handler ไม่ต้องรับรู้อะไร
			return nil , fiber.NewError(fiber.StatusNotFound , "customer not found")
		}
		logs.Error(err)
		// ส่ง error ไปยัง handler เพื่อให้ handler ไม่ต้องรับรู้อะไร
		return nil, fiber.NewError(fiber.StatusInternalServerError , "unexpected error")
	}

	cusRepo := CustomerResponse{
		CustomerID: customers.CustomerID,
		Name: customers.Name,
		Status: customers.Status,
	}

	return &cusRepo , nil 
}