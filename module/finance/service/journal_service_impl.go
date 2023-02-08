package service

import (
	"errors"
	"fmt"

	"github.com/riyan-eng/ngitung-dhuit/module/finance/controller/dto"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/repository"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/service/entity"
	"github.com/riyan-eng/ngitung-dhuit/util"
)

type journalServiceImpl struct {
	JournalRepository       repository.JournalRepository
	COARepository           repository.COARepository
	InventoryRepository     repository.InventoryRepository
	LinkedAccountRepository repository.LinkedAccountRepository
}

func NewJournalSerice(journalRepository repository.JournalRepository, coaRepository repository.COARepository, inventoryRepository repository.InventoryRepository, linkedAccountRepository repository.LinkedAccountRepository) JournalService {
	return &journalServiceImpl{
		JournalRepository:       journalRepository,
		COARepository:           coaRepository,
		InventoryRepository:     inventoryRepository,
		LinkedAccountRepository: linkedAccountRepository,
	}
}

func (service *journalServiceImpl) PurchaseJournal(dto *dto.PurchaseJournal) (err error) {
	// cek akun yang dicredit
	if dto.CreditAccount != util.COAAccountPayable && dto.CreditAccount != util.COACashInBank {
		return errors.New("coa can't be use")
	}

	if err := service.COARepository.GetByCode(dto.CreditAccount); err != nil {
		return errors.New("error getting coa credit")
	}

	// cek coa inventory code
	coaInventory, err := service.InventoryRepository.GetByCode(dto.InventoryCode)
	if err != nil {
		return errors.New("error getting coa inventory")
	}

	// jurnal menambah inventory
	journalInventory := entity.PurchaseJournal{
		Debet: entity.PurchaseJournalDebet{
			Code:   coaInventory,
			Amount: dto.Price * float64(dto.Quantity),
		},
		Credit: entity.PurchaseJournalCredit{
			Code:   dto.CreditAccount,
			Amount: dto.Price * float64(dto.Quantity),
		},
	}

	err = service.JournalRepository.PurchaseJournal(journalInventory)
	if err != nil {
		return errors.New("error journal add inventory")
	}

	// jurnal pajak
	if dto.PPNIncome == true {
		coaPPNIncome, err := service.LinkedAccountRepository.GetByCode("ppn_income")
		// journalInventory := entity.PurchaseJournal{
		// 	Debet: entity.PurchaseJournalDebet{
		// 		Code:   coaPPNIncome,
		// 		Amount: dto.Price * float64(dto.Quantity),
		// 	},
		// 	Credit: entity.PurchaseJournalCredit{
		// 		Code:   dto.CreditAccount,
		// 		Amount: dto.Price * float64(dto.Quantity),
		// 	},
		// }

		// err = service.JournalRepository.PurchaseJournal(journalInventory)
		// if err != nil {
		// 	return errors.New("error journal add inventory")
		// }
	}

	// var inventory float64 = dto.Price * float64(dto.Quantity)
	// var coaInventory string = coaInventorys

	var creditAmount float64 = inventory
	var creditCoa string = dto.CreditAccount

	fmt.Println(coaInventory)
	fmt.Println(creditAmount)
	fmt.Println(creditCoa)

	// creditAccount, err := repo.Journal.PurchaseJournal()

	// cek akun ppn masukan

	// var ppn_income_rate float64 = 8 / 100

	// debet
	// var inventory float64 = dto.Price * float64(dto.Quantity)

	// if dto.PPNIncome == true{
	// 	var ppn_income float64 = ppn_income_rate * inventory

	//jurnal
	// }

	// if dto.FreightPaid != 0{
	// 	var freight_paid float64 = dto.FreightPaid

	// jurnal
	// }

	// // kredit
	// cash_or_payable := dto.CreditAccount

	return nil
}

func (service *journalServiceImpl) SalesJournal() {

}
