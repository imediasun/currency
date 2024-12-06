package service

import (
	"app/business"
	"app/models"
)

type CurrencyExchangeService struct {
	exchangeLogic    business.ExchangeLogic
	transactionLogic business.TransactionLogic
}

// Инициализация сервиса
func NewCurrencyExchangeService(exchangeLogic business.ExchangeLogic, transactionLogic business.TransactionLogic) *CurrencyExchangeService {
	return &CurrencyExchangeService{
		exchangeLogic:    exchangeLogic,
		transactionLogic: transactionLogic,
	}
}

// Покупка валюты
func (s *CurrencyExchangeService) ExchangeCurrency(userID int, baseCurrency, targetCurrency string, amount float64) (*models.Transaction, error) {
	// Получение курса и расчёт конвертации
	result, err := s.exchangeLogic.CalculateExchange(baseCurrency, targetCurrency, amount)
	if err != nil {
		return nil, err
	}

	// Сохранение транзакции
	transaction, err := s.transactionLogic.CreateTransaction(userID, baseCurrency, targetCurrency, amount, result.ConvertedAmount, result.Fee)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Продажа валюты
func (s *CurrencyExchangeService) SellCurrency(userID int, sellCurrency, buyCurrency string, sellAmount float64) (*models.Transaction, error) {
	// Получение курса и расчёт конвертации
	result, err := s.exchangeLogic.CalculateExchange(sellCurrency, buyCurrency, sellAmount)
	if err != nil {
		return nil, err
	}

	// Сохранение транзакции
	transaction, err := s.transactionLogic.CreateTransaction(userID, sellCurrency, buyCurrency, sellAmount, result.ConvertedAmount, result.Fee)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
