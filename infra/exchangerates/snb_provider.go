package exchangerates

import (
	"encoding/json"
	"errors"
	"net/http"
)

type SNBExchangeRateProvider struct {
	APIEndpoint string
}

func NewSNBExchangeRateProvider(apiEndpoint string) *SNBExchangeRateProvider {
	return &SNBExchangeRateProvider{APIEndpoint: apiEndpoint}
}

func (p *SNBExchangeRateProvider) GetRate(baseCurrency, targetCurrency string) (float64, error) {
	resp, err := http.Get(p.APIEndpoint)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	rates, ok := data["rates"].(map[string]interface{})
	if !ok {
		return 0, errors.New("invalid rates format")
	}

	rate, ok := rates[targetCurrency].(float64)
	if !ok {
		return 0, errors.New("rate not found for target currency")
	}

	return rate, nil
}
