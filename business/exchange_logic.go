package business

import "app/models"

type ExchangeLogic interface {
	CalculateExchange(baseCurrency, targetCurrency string, amount float64) (*models.ExchangeResult, error)
}

type exchangeLogicImpl struct {
	rateProvider RateProvider
	feePolicy    FeePolicy
}

func NewExchangeLogic(rateProvider RateProvider, feePolicy FeePolicy) ExchangeLogic {
	return &exchangeLogicImpl{
		rateProvider: rateProvider,
		feePolicy:    feePolicy,
	}
}

func (e *exchangeLogicImpl) CalculateExchange(baseCurrency, targetCurrency string, amount float64) (*models.ExchangeResult, error) {
	// Получение курса
	rate, err := e.rateProvider.GetRate(baseCurrency, targetCurrency)
	if err != nil {
		return nil, err
	}

	// Расчёт суммы и комиссии
	fee := e.feePolicy.Calculate(amount)
	convertedAmount := amount*rate - fee

	return &models.ExchangeResult{
		Rate:            rate,
		ConvertedAmount: convertedAmount,
		Fee:             fee,
	}, nil
}
