package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/controller/dto"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/service"
	"github.com/riyan-eng/ngitung-dhuit/util"
)

type journalService struct {
	Journal service.JournalService
}

func NewJournalController(service service.JournalService, route *fiber.App) {
	srvc := &journalService{
		Journal: service,
	}

	journalRoute := route.Group("/journal")
	journalRoute.Post("/purchase", srvc.PurchaseJournal)
	journalRoute.Post("/sales", srvc.SalesJournal)
}

func (service journalService) PurchaseJournal(c *fiber.Ctx) error {
	body := new(dto.PurchaseJournal)

	// parsing body
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "bad",
		})
	}

	// valdate body
	if err := util.Validate(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "bad",
		})
	}

	// communicate service
	ctx := c.Context()
	if err := service.Journal.PurchaseJournal(ctx, body); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "bad",
		})
	}

	// return
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}

func (service journalService) SalesJournal(c *fiber.Ctx) error {
	return nil
}
