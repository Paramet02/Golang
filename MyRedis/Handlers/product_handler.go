package handlers

import (
	service "github.com/paramet02/Service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type productHandler struct {
	pService service.ProductService
}

// NewproductHandler สร้างและคืนค่า handler สำหรับการจัดการกับสินค้าผ่าน service
func NewproductHandler(pService service.ProductService) ProductHandler {
	return productHandler{pService}
}

// CreateProduct ฟังก์ชันสำหรับการสร้างสินค้าผ่าน HTTP POST request
func (p productHandler) CreateProduct(c *fiber.Ctx) error {

	var product service.Product

	// ใช้ BodyParser เพื่อแปลง body ของ request เป็น struct product
	if err := c.BodyParser(&product); err != nil {
        return c.SendStatus(fiber.StatusBadRequest) // ถ้าเกิดข้อผิดพลาดในการแปลงข้อมูล
    }

	// เรียกใช้ service เพื่อสร้างสินค้า
	err := p.pService.CreateProduct(product)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest) // ถ้าเกิดข้อผิดพลาดในการสร้างสินค้า
	}

	// ถ้าสร้างสำเร็จ ส่ง response กลับไป
	return c.JSON(fiber.Map{
        "status":  "success",
        "message": "Product added successfully",
    })
}

// UpdateProduct ฟังก์ชันสำหรับการอัปเดตข้อมูลสินค้าผ่าน HTTP PUT request
func (p productHandler) UpdateProduct(c *fiber.Ctx) error {
	// รับค่า id จาก URL params และแปลงเป็น int
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid ID format", // ถ้า id ไม่ใช่ตัวเลขที่ถูกต้อง
		})
	}

	// รับค่า product จาก body ของ request
	var product service.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request body", // ถ้าไม่สามารถแปลง body เป็น struct ได้
		})
	}

	// ตรวจสอบว่า product.Name และ product.Quantity ถูกต้อง
	if product.Name == "" || product.Quantity <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product data", // ถ้าข้อมูลสินค้าไม่ครบถ้วน
		})
	}

	// ตั้งค่า ID ของสินค้าให้เป็น id ที่รับมาจาก URL
	product.ID = id

	// เรียกใช้ service เพื่ออัปเดตข้อมูลสินค้า
	err = p.pService.UpdateProduct(product)
	if err != nil {
		// ส่งข้อผิดพลาดจาก service กลับไป
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(), // ข้อความที่ได้จาก service
		})
	}

	// ถ้าอัปเดตสำเร็จ ส่ง response กลับไป
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product updated successfully",
	})
}

// GetProduct ฟังก์ชันสำหรับการค้นหาสินค้าจากชื่อผ่าน HTTP GET request
func (p productHandler) GetProduct(c *fiber.Ctx) error {
    // ดึงค่า "name" จาก query string
    name := c.Query("name") 

    // ถ้าไม่มีชื่อสินค้าผ่าน query string มา ให้ส่งสถานะ 400
    if name == "" {
        return c.SendStatus(fiber.StatusBadRequest)
    }

    // เรียก service เพื่อค้นหาสินค้าจากชื่อ
    products, err := p.pService.GetProduct(name)
    if err != nil {
        return err // ส่งข้อผิดพลาดจาก service ถ้ามี
    }

    // ส่งผลลัพธ์กลับไปในรูปแบบ JSON
    return c.JSON(fiber.Map{
        "status":  "ok",
        "products": products, // ส่งสินค้าในรูปแบบ JSON
    })
}

// GetProducts ฟังก์ชันสำหรับการดึงข้อมูลสินค้าทั้งหมดผ่าน HTTP GET request
func (p productHandler) GetProducts(c *fiber.Ctx) error {
    // เรียกใช้ service เพื่อดึงข้อมูลสินค้าทั้งหมด
    produnct , err := p.pService.GetProducts()
	if err != nil {
		return err // ถ้ามีข้อผิดพลาดจาก service ให้ส่งกลับ
	}

	// ส่งผลลัพธ์กลับไปในรูปแบบ JSON
	response := fiber.Map{
		"status": "ok",
		"product": produnct, // ส่งรายการสินค้าทั้งหมด
	} 
	return c.JSON(response)
}
