package handlers

import (
	service "github.com/paramet02/Service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type productHandler struct {
	pService service.ProductService
}

func NewproductHandler(pService service.ProductService) ProductHandler {
	return productHandler{pService}
}

func (p productHandler) CreateProduct(c *fiber.Ctx) error {

	var product service.Product

	if err := c.BodyParser(&product); err != nil {
        return c.SendStatus(fiber.StatusBadRequest)
    }

	err := p.pService.CreateProduct(product)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
        "status":  "success",
        "message": "Product added successfully",
    })
}

func (p productHandler) UpdateProduct(c *fiber.Ctx) error {
	// รับค่า id จาก URL
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid ID format",
		})
	}

	// รับค่า product จาก body
	var product service.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request body",
		})
	}

	// ตรวจสอบให้แน่ใจว่า product.Name และ product.Quantity มีค่า
	if product.Name == "" || product.Quantity <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product data",
		})
	}

	// ตั้งค่า ID ที่รับมาจาก URL ไปยัง Product ที่ส่งมา
	product.ID = id

	// เรียกใช้ service เพื่ออัปเดต
	err = p.pService.UpdateProduct(product)
	if err != nil {
		// แสดงข้อความข้อผิดพลาดที่มาจาก service
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product updated successfully",
	})
}

func (p productHandler) GetProduct(c *fiber.Ctx) error {
    // ดึงค่า "name" จาก query string
    name := c.Query("name") 

    if name == "" {
        return c.SendStatus(fiber.StatusBadRequest) // ส่งสถานะ 400 ถ้าไม่มีชื่อสินค้าใน query string
    }

    // เรียก service เพื่อค้นหาสินค้า
    products, err := p.pService.GetProduct(name)
    if err != nil {
        return err
    }

    // ส่งผลลัพธ์กลับในรูปแบบ JSON
    return c.JSON(fiber.Map{
        "status":  "ok",
        "products": products,
    })
}

func (p productHandler) GetProducts(c *fiber.Ctx) error {
    produnct , err := p.pService.GetProducts()
	if err != nil {
		return err
	}
	response := fiber.Map{
		"status": "ok",
		"product": produnct,
	} 
	return c.JSON(response)
}