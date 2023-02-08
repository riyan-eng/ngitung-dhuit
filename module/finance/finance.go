package finance

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/ngitung-dhuit/config"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/controller"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/repository"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/service"
)

func Setup(app *fiber.App) {
	journalRepository := repository.NewJournalRepository(config.DB)
	coaRepository := repository.NewCOARepository(config.DB)
	inventoryRepository := repository.NewInventoryRepository(config.DB)
	linkedAccountRepository := repository.NewLinkedAccountRepository(config.DB)
	taxRepository := repository.NewTaxRepository(config.DB)
	srvc := service.NewJournalSerice(journalRepository, coaRepository, inventoryRepository, linkedAccountRepository, taxRepository)
	controller.NewJournalController(srvc, app)
}
