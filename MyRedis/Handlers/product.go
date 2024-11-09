package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface{
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	GetProducts(c *fiber.Ctx) error
	GetProduct(c *fiber.Ctx) error
} 