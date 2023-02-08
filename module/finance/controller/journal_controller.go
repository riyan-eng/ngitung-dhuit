package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/service"
)

type journalService struct {
	Journal service.JournalService
}

func NewJournalController(service service.JournalService, route *fiber.App) {
	s := &journalService{
		Journal: service,
	}

	journalRoute := route.Group("/journal")
	journalRoute.Post("/purchase", s.PurchaseJournal)
	journalRoute.Post("/sales", s.SalesJournal)
}

func (service journalService) PurchaseJournal(c *fiber.Ctx) error {
	return nil
}

func (service journalService) SalesJournal(c *fiber.Ctx) error {
	return nil
}
