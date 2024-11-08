package handler

import (
	"paramet/service"
	"github.com/gofiber/fiber/v2"
)

type catalogHandler struct {
	catalogService service.CatalogService
}

func NewcatalogHandler(catalogService service.CatalogService) CatalogHandler {
	return catalogHandler{catalogService}
}

func (h catalogHandler) GetProduct(c *fiber.Ctx) error {
	produnct , err := h.catalogService.GetProduct()
	if err != nil {
		return err
	}
	response := fiber.Map{
		"status": "ok",
		"product": produnct,
	} 
	return c.JSON(response)
}