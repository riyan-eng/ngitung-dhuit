package finance

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/ngitung-dhuit/config"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/controller"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/repository"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/service"
)

func Setup(app *fiber.App) {
	repo := repository.NewJournalRepository(config.DB)
	srvc := service.NewJournalSerice(repo)
	controller.NewJournalController(srvc, app)
}
