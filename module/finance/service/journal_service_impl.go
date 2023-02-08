package service

import (
	"errors"

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
	TaxRepository           repository.TaxRepository
}

func NewJournalSerice(journalRepository repository.JournalRepository, coaRepository repository.COARepository, inventoryRepository repository.InventoryRepository, linkedAccountRepository repository.LinkedAccountRepository, taxRepository repository.TaxRepository) JournalService {
	return &journalServiceImpl{
		JournalRepository:       journalRepository,
		COARepository:           coaRepository,
		InventoryRepository:     inventoryRepository,
		LinkedAccountRepository: linkedAccountRepository,
		TaxRepository:           taxRepository,
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
	var inventoryAmount float64 = dto.Price * float64(dto.Quantity)

	journalInventory := entity.PurchaseJournal{
		Debet: entity.PurchaseJournalDebet{
			Code:   coaInventory,
			Amount: inventoryAmount,
		},
		Credit: entity.PurchaseJournalCredit{
			Code:   dto.CreditAccount,
			Amount: inventoryAmount,
		},
	}

	err = service.JournalRepository.PurchaseJournal(journalInventory)
	if err != nil {
		return errors.New("error journal add inventory")
	}

	// jurnal pajak
	if dto.PPNIncome == true {
		coaPPNIncome, err := service.LinkedAccountRepository.GetByCode("ppn_income")
		if err != nil {
			return errors.New("error getting linked account")
		}
		tax, err := service.TaxRepository.GetByCoa(coaPPNIncome)
		if err != nil {
			return errors.New("error getting tax rate")
		}

		var taxRate float64 = float64(tax) / 100
		var taxAmount float64 = inventoryAmount * taxRate

		journalPPNIncome := entity.PurchaseJournal{
			Debet: entity.PurchaseJournalDebet{
				Code:   coaPPNIncome,
				Amount: taxAmount,
			},
			Credit: entity.PurchaseJournalCredit{
				Code:   dto.CreditAccount,
				Amount: taxAmount,
			},
		}

		err = service.JournalRepository.PurchaseJournal(journalPPNIncome)
		if err != nil {
			return errors.New("error journal add inventory")
		}
	}

	// jurnal biaya angkut
	if dto.FreightPaid > 0 {
		coaFreightPaid, err := service.LinkedAccountRepository.GetByCode("freight_paid")
		if err != nil {
			return errors.New("error getting linked account")
		}

		journalFreightPaid := entity.PurchaseJournal{
			Debet: entity.PurchaseJournalDebet{
				Code:   coaFreightPaid,
				Amount: dto.FreightPaid,
			},
			Credit: entity.PurchaseJournalCredit{
				Code:   dto.CreditAccount,
				Amount: dto.FreightPaid,
			},
		}

		err = service.JournalRepository.PurchaseJournal(journalFreightPaid)
		if err != nil {
			return errors.New("error journal add inventory")
		}
	}

	return nil
}

func (service *journalServiceImpl) SalesJournal() {

}
