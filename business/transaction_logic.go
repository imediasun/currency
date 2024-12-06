package business

import "app/models"

type TransactionLogic interface {
	CreateTransaction(userID int, baseCurrency, targetCurrency string, amount, convertedAmount, fee float64) (*models.Transaction, error)
}

type transactionLogicImpl struct {
	transactionRepo TransactionRepository
}

func NewTransactionLogic(transactionRepo TransactionRepository) TransactionLogic {
	return &transactionLogicImpl{
		transactionRepo: transactionRepo,
	}
}

func (t *transactionLogicImpl) CreateTransaction(userID int, baseCurrency, targetCurrency string, amount, convertedAmount, fee float64) (*models.Transaction, error) {
	transaction := &models.Transaction{
		UserID:          userID,
		BaseCurrency:    baseCurrency,
		TargetCurrency:  targetCurrency,
		Amount:          amount,
		ConvertedAmount: convertedAmount,
		Fee:             fee,
	}

	err := t.transactionRepo.SaveTransaction(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
