package models

type Transaction struct {
	UserID          int
	BaseCurrency    string
	TargetCurrency  string
	Amount          float64
	ConvertedAmount float64
	Fee             float64
}

type ExchangeResult struct {
	Rate            float64
	ConvertedAmount float64
	Fee             float64
}
