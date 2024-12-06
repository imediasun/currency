package data

import (
	"app/models"
	"database/sql"
)

type TransactionRepository interface {
	SaveTransaction(transaction *models.Transaction) error
}

type transactionRepoImpl struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepoImpl{db: db}
}

func (r *transactionRepoImpl) SaveTransaction(transaction *models.Transaction) error {
	query := `INSERT INTO transactions (user_id, base_currency, target_currency, amount, converted_amount, fee) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, transaction.UserID, transaction.BaseCurrency, transaction.TargetCurrency, transaction.Amount, transaction.ConvertedAmount, transaction.Fee)
	return err
}
