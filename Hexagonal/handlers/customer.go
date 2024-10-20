package handler

// Handler Layer
// It is the layer that handles receiving requests and sending responses to users.
// It calls the service layer through various handlers such as GetCustomers and GetCustomer.

import (
	"paramet/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type customerHandler struct {
	cs service.CustomerService
}
// return customerHandler ที่เก็บ service ไว้ ใน struct 
// เพื่อ ให้ handler สามารถทำงานได้โดยไม่ต้องรู้รายละเอียดว่าข้อมูลมาจากไหนหรือฐานข้อมูลทำงานยังไง
func NewCustomerHandler(cs service.CustomerService) customerHandler {
	return customerHandler{cs : cs}
}
// GetCustomer(request and response) = fiber.ctx
func(h customerHandler) GetCustomers(c *fiber.Ctx) error {
	customers, err := h.cs.GetCustomers()
	if err != nil {
		return err
	}
	return c.JSON(customers)
}

func(h *customerHandler) GetCustomer(c *fiber.Ctx) error {
	// main.go customer/:id == c.Params("id")
	CustomerID ,err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	customer , err := h.cs.GetCustomer(CustomerID)
	if err != nil {
		return err
	}

	return c.JSON(customer)
}

