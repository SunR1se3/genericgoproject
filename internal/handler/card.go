package handler

import (
	"GenericProject/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) CreateCard(c *fiber.Ctx) error {
	formData := new(domain.Card)
	err := c.BodyParser(formData)
	if err != nil {
		return c.JSON(err)
	}
	id, err := h.services.Card.CreateCard(*formData)
	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON(fiber.Map{
		"Id": id,
	})
}

func (h *Handler) GetCardById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.JSON(err)
	}
	data, err := h.services.Card.GetById(id)
	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON(fiber.Map{
		"data": data,
	})
}

func (h *Handler) UpdateCard(c *fiber.Ctx) error {
	formData := new(domain.Card)
	err := c.BodyParser(formData)
	if err != nil {
		return c.JSON(err)
	}
	err = h.services.Card.UpdateCard(*formData)
	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON(fiber.Map{
		"Status": true,
	})
}
