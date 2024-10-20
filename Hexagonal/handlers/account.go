package handler

import (
	"paramet/service"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type accountHandler struct {
	accSrv service.AccountService 
}

func NewacccountHandler(accSrv service.AccountService) *accountHandler {
	return &accountHandler{accSrv : accSrv}
}

// NewAccount handles account creation requests.
func (h *accountHandler) NewAccount(c *fiber.Ctx) error {
	customerID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// คือหนึ่งในค่า Content-Type ที่ระบุว่าข้อมูลที่ส่งมานั้นเป็น JSON (JavaScript Object Notation) 
	// ซึ่งเป็นรูปแบบมาตรฐานในการแลกเปลี่ยนข้อมูลที่มีโครงสร้างระหว่าง client และ server
	if c.Get("Content-Type") != "application/json" {
		return fiber.NewError(fiber.StatusBadRequest, "Content-Type must be application/json")
	}

	request := service.NewAccountRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	response, err := h.accSrv.NewAccount(customerID, request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// GetAccounts handles the retrieval of accounts for a customer.
func (h *accountHandler) GetAccount(c *fiber.Ctx) error {
	customerID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	responses, err := h.accSrv.GetAccount(customerID)
	if err != nil {
		return err
	}

	return c.JSON(responses)
}