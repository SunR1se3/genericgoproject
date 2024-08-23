package handler

import (
	"GenericProject/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init(app *fiber.App) {
	api := app.Group("api")

	card := api.Group("card")
	card.Post("/", h.CreateCard)
	card.Get("/:id", h.GetCardById)
	card.Get("/", h.GetAllCards)
	card.Put("/", h.UpdateCard)
	card.Delete("/:id", h.DeleteCard)
}
