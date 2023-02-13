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
	TransactionRepository   repository.TransactionRepository
	JournalRepository       repository.JournalRepository
	COARepository           repository.COARepository
	InventoryRepository     repository.InventoryRepository
	LinkedAccountRepository repository.LinkedAccountRepository
	TaxRepository           repository.TaxRepository
	SubsidiaryLedger        repository.SubsidiaryLedgerRepository
}

func NewJournalSerice(transactionRepository repository.TransactionRepository, journalRepository repository.JournalRepository, coaRepository repository.COARepository, inventoryRepository repository.InventoryRepository, linkedAccountRepository repository.LinkedAccountRepository, taxRepository repository.TaxRepository, subsidiaryLedger repository.SubsidiaryLedgerRepository) JournalService {
	return &journalServiceImpl{
		TransactionRepository:   transactionRepository,
		JournalRepository:       journalRepository,
		COARepository:           coaRepository,
		InventoryRepository:     inventoryRepository,
		LinkedAccountRepository: linkedAccountRepository,
		TaxRepository:           taxRepository,
		SubsidiaryLedger:        subsidiaryLedger,
	}
}

func (service *journalServiceImpl) PurchaseJournal(dto *dto.PurchaseJournal) (err error) {
	// jumlah transaksi
	// var inventoryAmount float64 = dto.Price * float64(dto.Quantity)

	// cek akun yang dicredit
	if dto.CreditAccount != util.COAAccountPayable && dto.CreditAccount != util.COACashInBank {
		return errors.New("coa can't be use")
	}

	if err := service.COARepository.GetByCode(dto.CreditAccount); err != nil {
		return errors.New("error getting coa credit")
	}

	// // cek coa inventory code
	// _, err = service.InventoryRepository.GetByCode(dto.InventoryCode)
	// if err != nil {
	// 	return errors.New("error getting coa inventory")
	// }

	// // get saldo terakhir
	// currentBalance, err := service.InventoryRepository.CurrentBalance(dto.InventoryCode)
	// if err != nil {
	// 	return err
	// }

	// // kalkulasi saldo terbaru
	// balanceQuantity := currentBalance.Quantity + int(dto.Quantity)
	// balanceAmount := currentBalance.Amount + inventoryAmount
	// balancePrice := balanceAmount / float64(balanceQuantity)

	// // entity tambah
	// inventoryEntity := entity.InventoryIn{
	// 	Quantity:        int(dto.Quantity),
	// 	Price:           dto.Price,
	// 	Amount:          inventoryAmount,
	// 	BalanceQuantity: balanceQuantity,
	// 	BalancePrice:    balancePrice,
	// 	BalanceAmount:   balanceAmount,
	// }

	// if err := service.InventoryRepository.In(dto.InventoryCode, inventoryEntity); err != nil {
	// 	return err
	// }

	// // masukkan ke tabel transaksi
	// transaction_id, err := service.TransactionRepository.Insert(dto.Description, inventoryAmount)
	// if err != nil {
	// 	return errors.New("error insert transaction")
	// }

	// // jurnal menambah inventory

	// journalInventory := entity.PurchaseJournal{
	// 	TransactionID: transaction_id,
	// 	Debet: entity.PurchaseJournalDebet{
	// 		Code:   util.COAMerchandiseInventory,
	// 		Amount: inventoryAmount,
	// 	},
	// 	Credit: entity.PurchaseJournalCredit{
	// 		Code:   dto.CreditAccount,
	// 		Amount: inventoryAmount,
	// 	},
	// }

	// err = service.JournalRepository.PurchaseJournal(journalInventory)
	// if err != nil {
	// 	return errors.New("error journal add inventory")
	// }

	// // jurnal pajak
	// var taxAmount float64
	// if dto.PPNIncome {
	// 	coaPPNIncome, err := service.LinkedAccountRepository.GetByCode("ppn_income")
	// 	if err != nil {
	// 		return errors.New("error getting linked account")
	// 	}
	// 	tax, err := service.TaxRepository.GetByCoa(coaPPNIncome)
	// 	if err != nil {
	// 		return errors.New("error getting tax rate")
	// 	}

	// 	var taxRate float64 = float64(tax) / 100
	// 	taxAmount = inventoryAmount * taxRate

	// 	journalPPNIncome := entity.PurchaseJournal{
	// 		TransactionID: transaction_id,
	// 		Debet: entity.PurchaseJournalDebet{
	// 			Code:   coaPPNIncome,
	// 			Amount: taxAmount,
	// 		},
	// 		Credit: entity.PurchaseJournalCredit{
	// 			Code:   dto.CreditAccount,
	// 			Amount: taxAmount,
	// 		},
	// 	}

	// 	err = service.JournalRepository.PurchaseJournal(journalPPNIncome)
	// 	if err != nil {
	// 		return errors.New("error journal add inventory")
	// 	}
	// }

	// // jurnal biaya angkut
	// if dto.FreightPaid > 0 {
	// 	coaFreightPaid, err := service.LinkedAccountRepository.GetByCode("freight_paid")
	// 	if err != nil {
	// 		return errors.New("error getting linked account")
	// 	}

	// 	journalFreightPaid := entity.PurchaseJournal{
	// 		TransactionID: transaction_id,
	// 		Debet: entity.PurchaseJournalDebet{
	// 			Code:   coaFreightPaid,
	// 			Amount: dto.FreightPaid,
	// 		},
	// 		Credit: entity.PurchaseJournalCredit{
	// 			Code:   dto.CreditAccount,
	// 			Amount: dto.FreightPaid,
	// 		},
	// 	}

	// 	err = service.JournalRepository.PurchaseJournal(journalFreightPaid)
	// 	if err != nil {
	// 		return errors.New("error journal add inventory")
	// 	}
	// }

	// if dto.CreditAccount == util.COAAccountPayable && dto.SupplierCode != "" {
	// 	// insert ke buku besar pembantu utang
	// 	if err := service.SubsidiaryLedger.InsertPayable(dto.SupplierCode, inventoryAmount+taxAmount+dto.FreightPaid); err != nil {
	// 		return err
	// 	}
	// }

	// fmt.Println(dto)
	var inventoryAmount float64
	var inventoryDiscount float64
	for _, good := range dto.Goods {
		amount := good.Price * float64(good.Qty)
		discount := amount * float64(good.Discount) / 100
		inventoryAmount += amount
		inventoryDiscount += discount

		// fmt.Println("good_code:", good.GoodCode)
		// fmt.Println("quantity:", good.Qty)
		// fmt.Println("price:", good.Price)
		// fmt.Println("amount:", int(amount))

		// cek coa inventory code
		_, err = service.InventoryRepository.GetByCode(good.GoodCode)
		if err != nil {
			return errors.New("error getting coa inventory")
		}

		// get saldo terakhir
		currentBalance, err := service.InventoryRepository.CurrentBalance(good.GoodCode)
		if err != nil {
			return err
		}

		// kalkulasi saldo terbaru
		balanceQuantity := currentBalance.Quantity + good.Qty
		balanceAmount := currentBalance.Amount + amount
		balancePrice := balanceAmount / float64(balanceQuantity)

		// entity tambah
		inventoryEntity := entity.InventoryIn{
			Quantity:        good.Qty,
			Price:           good.Price,
			Amount:          amount,
			BalanceQuantity: balanceQuantity,
			BalancePrice:    balancePrice,
			BalanceAmount:   balanceAmount,
		}

		if err := service.InventoryRepository.In(good.GoodCode, inventoryEntity); err != nil {
			return err
		}
	}

	fmt.Println("total:", int(inventoryAmount))
	fmt.Println("diskon:", int(inventoryDiscount))
	fmt.Println("freight paid:", int(dto.FreightPaid))
	fmt.Println("ppn income:", dto.PPNIncome)
	fmt.Println("credit account:", dto.CreditAccount)
	fmt.Println("supplier code:", dto.SupplierCode)
	fmt.Println("description:", dto.Description)

	inventoryTotal := inventoryAmount - inventoryDiscount

	// jurnal pajak
	var ppnAmount float64
	if dto.PPNIncome {
		coaPPNIncome, err := service.LinkedAccountRepository.GetByCode("ppn_income")
		if err != nil {
			return errors.New("error getting linked account")
		}
		tax, err := service.TaxRepository.GetByCoa(coaPPNIncome)
		if err != nil {
			return errors.New("error getting tax rate")
		}

		var taxRate float64 = float64(tax) / 100
		ppnAmount = inventoryAmount * taxRate

		journalPPNIncome := entity.PurchaseJournal{
			TransactionID: transaction_id,
			Debet: entity.PurchaseJournalDebet{
				Code:   coaPPNIncome,
				Amount: ppnAmount,
			},
			Credit: entity.PurchaseJournalCredit{
				Code:   dto.CreditAccount,
				Amount: ppnAmount,
			},
		}
		err = service.JournalRepository.PurchaseJournal(journalPPNIncome)
		if err != nil {
			return errors.New("error journal add inventory")
		}
	}
	// ppnAmount := 1
	total := inventoryTotal + ppnAmount + dto.FreightPaid
	fmt.Println(int(total))
	return nil
}

// func insertTableTransaction(transactionId, description string, amount float64) error {
// 	return nil
// }

// func insertTableInventory(inventoryCode, transactionId string, inventoryEntity entity.Inventory) error {
// 	return nil
// }

// func insertTablePurchaseJournal(purchaseJournal entity.PurchaseJournal) error {
// 	return nil
// }

func (service *journalServiceImpl) SalesJournal() {

}
