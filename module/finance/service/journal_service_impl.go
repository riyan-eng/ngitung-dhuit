package service

import (
	"errors"

	"github.com/riyan-eng/ngitung-dhuit/module/finance/controller/dto"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/repository"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/service/entity"
	"github.com/riyan-eng/ngitung-dhuit/util"
	"github.com/valyala/fasthttp"
)

type journalServiceImpl struct {
	TransactionRepository   repository.TransactionRepository
	JournalRepository       repository.JournalRepository
	COARepository           repository.COARepository
	InventoryRepository     repository.InventoryRepository
	LinkedAccountRepository repository.LinkedAccountRepository
	TaxRepository           repository.TaxRepository
	SubsidiaryLedger        repository.SubsidiaryLedgerRepository
	SupplierRepository      repository.SupplierRepository
}

func NewJournalSerice(transactionRepository repository.TransactionRepository, journalRepository repository.JournalRepository, coaRepository repository.COARepository, inventoryRepository repository.InventoryRepository, linkedAccountRepository repository.LinkedAccountRepository, taxRepository repository.TaxRepository, subsidiaryLedger repository.SubsidiaryLedgerRepository, supplierRepository repository.SupplierRepository) JournalService {
	return &journalServiceImpl{
		TransactionRepository:   transactionRepository,
		JournalRepository:       journalRepository,
		COARepository:           coaRepository,
		InventoryRepository:     inventoryRepository,
		LinkedAccountRepository: linkedAccountRepository,
		TaxRepository:           taxRepository,
		SubsidiaryLedger:        subsidiaryLedger,
		SupplierRepository:      supplierRepository,
	}
}

func (service *journalServiceImpl) PurchaseJournal(ctx *fasthttp.RequestCtx, dto *dto.PurchaseJournal) (err error) {
	// cek akun yang dicredit
	if dto.CreditAccount != util.COAAccountPayable && dto.CreditAccount != util.COACashInBank {
		return errors.New("coa can't be use")
	}
	if err := service.COARepository.FindOneByCode(ctx, dto.CreditAccount); err != nil {
		return errors.New("error getting coa credit")
	}

	// cek supplier
	if err := service.SupplierRepository.FindOneByCode(ctx, dto.SupplierCode); err != nil {
		return err
	}

	// cek inventory
	var inventoryAmount float64
	var inventoryDiscount float64
	for _, good := range dto.Goods {
		// cek coa inventory code
		_, err = service.InventoryRepository.FindOneByCode(ctx, good.GoodCode)
		if err != nil {
			return errors.New("error getting coa inventory")
		}
		amount := good.Price * float64(good.Qty)
		discount := amount * float64(good.Discount) / 100
		inventoryAmount += amount
		inventoryDiscount += discount
	}

	inventoryTotal := inventoryAmount - inventoryDiscount

	// cek pajak
	var ppnAmount float64
	var coaPPNIncome string
	if dto.PPNIncome {
		coaPPNIncome, err = service.LinkedAccountRepository.FindOneByCode(ctx, "ppn_income")
		if err != nil {
			return errors.New("error getting linked account")
		}
		tax, err := service.TaxRepository.FindOneByCoa(ctx, coaPPNIncome)
		if err != nil {
			return errors.New("error getting tax rate")
		}
		var taxRate float64 = float64(tax) / 100
		ppnAmount = inventoryAmount * taxRate
	}

	// cek freight paid
	coaFreightPaid, err := service.LinkedAccountRepository.FindOneByCode(ctx, "freight_paid")
	if err != nil {
		return errors.New("error getting linked account")
	}

	total := inventoryTotal + ppnAmount + dto.FreightPaid

	// insert transaction
	var transactionID string
	if transactionID, err = service.TransactionRepository.InsertOne(ctx, dto.Description, total); err != nil {
		return err
	}

	// jurnal menambah inventory
	journalInventory := entity.PurchaseJournal{
		TransactionID: transactionID,
		Debet: entity.PurchaseJournalDebet{
			Code:   util.COAMerchandiseInventory,
			Amount: inventoryAmount,
		},
		Credit: entity.PurchaseJournalCredit{
			Code:   dto.CreditAccount,
			Amount: inventoryAmount,
		},
	}
	if err := service.JournalRepository.InsertOnePurchaseJournal(ctx, journalInventory); err != nil {
		return errors.New("error journal add inventory")
	}

	// jurnal diskon
	if inventoryDiscount > 0 {
		journalInventory := entity.PurchaseJournal{
			TransactionID: transactionID,
			Debet: entity.PurchaseJournalDebet{
				Code:   dto.CreditAccount,
				Amount: inventoryDiscount,
			},
			Credit: entity.PurchaseJournalCredit{
				Code:   util.COAPurchaseDiscount,
				Amount: inventoryDiscount,
			},
		}
		if err := service.JournalRepository.InsertOnePurchaseJournal(ctx, journalInventory); err != nil {
			return errors.New("error journal add inventory")
		}
	}

	// insert inventory
	for _, good := range dto.Goods {
		amount := good.Price * float64(good.Qty)
		// get saldo terakhir
		currentBalance, err := service.InventoryRepository.CurrentBalance(ctx, good.GoodCode)
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
		if err := service.InventoryRepository.In(ctx, transactionID, good.GoodCode, inventoryEntity); err != nil {
			return err
		}
	}

	// insert tax
	if ppnAmount > 0 {
		journalPPNIncome := entity.PurchaseJournal{
			TransactionID: transactionID,
			Debet: entity.PurchaseJournalDebet{
				Code:   coaPPNIncome,
				Amount: ppnAmount,
			},
			Credit: entity.PurchaseJournalCredit{
				Code:   dto.CreditAccount,
				Amount: ppnAmount,
			},
		}
		if err := service.JournalRepository.InsertOnePurchaseJournal(ctx, journalPPNIncome); err != nil {
			return errors.New("error journal add inventory")
		}
	}

	// insert freight paid
	if dto.FreightPaid > 0 {
		journalFreightPaid := entity.PurchaseJournal{
			TransactionID: transactionID,
			Debet: entity.PurchaseJournalDebet{
				Code:   coaFreightPaid,
				Amount: dto.FreightPaid,
			},
			Credit: entity.PurchaseJournalCredit{
				Code:   dto.CreditAccount,
				Amount: dto.FreightPaid,
			},
		}
		err = service.JournalRepository.InsertOnePurchaseJournal(ctx, journalFreightPaid)
		if err != nil {
			return errors.New("error journal add inventory")
		}
	}

	// insert buku besar pembantu utang
	if dto.CreditAccount == util.COAAccountPayable && dto.SupplierCode != "" {
		if err := service.SubsidiaryLedger.InsertOnePayable(ctx, dto.SupplierCode, transactionID, total); err != nil {
			return err
		}
	}
	return err
}

func (service *journalServiceImpl) SalesJournal() {

}

func (service *journalServiceImpl) CashPaymentJournal() {
	//
}
