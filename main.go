package main

import (
	"app/business"
	"app/data"
	"app/infrastructure/exchangerates"
	"app/service"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite драйвер
)

const (
	dbPath         = "./currency_exchange.db"
	migrationsPath = "./migrations"
)

func main() {
	// 1. Инициализация базы данных
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 2. Выполнение миграций
	err = data.RunMigrations(db, migrationsPath)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// 3. Создание инфраструктуры
	apiEndpoint := "https://data.snb.ch/api/some_endpoint"
	rateProvider := exchangerates.NewSNBExchangeRateProvider(apiEndpoint)

	// 4. Определение политики комиссий
	feePolicy := business.NewFixedFeePolicy(1.5) // 1.5% комиссия

	// 5. Инициализация бизнес-логики
	exchangeLogic := business.NewExchangeLogic(rateProvider, feePolicy)
	transactionRepo := data.NewTransactionRepository(db)
	transactionLogic := business.NewTransactionLogic(transactionRepo)

	// 6. Создание сервиса
	currencyService := service.NewCurrencyExchangeService(exchangeLogic, transactionLogic)

	// 7. Выполнение операций

	// Пример: Покупка валюты (100 USD -> EUR)
	transaction, err := currencyService.ExchangeCurrency(1, "USD", "EUR", 100.0)
	if err != nil {
		log.Fatalf("Error during exchange: %v", err)
	}
	fmt.Printf("Transaction (Buy): %+v\n", transaction)

	// Пример: Продажа валюты (125 EUR -> USD)
	transaction, err = currencyService.SellCurrency(1, "EUR", "USD", 125.0)
	if err != nil {
		log.Fatalf("Error during sell: %v", err)
	}
	fmt.Printf("Transaction (Sell): %+v\n", transaction)
}
