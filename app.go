package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
)

var dbCon *sql.DB
var err error
var totalSpending float64

// App struct
type App struct {
	ctx context.Context
}

type Transaction struct {
	Date        time.Time
	Description string
	Category    string
	Amount      float64
}

type TotalPerStore struct {
	Name              string
	Amount            float64
	TrancationsNumber int
}
type Category struct {
	Category       string
	PlaceBreakdown []TotalPerStore
	Amount         float64
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	dbCon, err = startDb()
	if err != nil {
		log.Fatal(err)
	}
}

func (b *App) shutdown(ctx context.Context) {
	if dbCon != nil {
		dbCon.Close()
	}
}

func (a *App) ParseCsv(content string) error {
	reader := csv.NewReader(strings.NewReader(content))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse CSV: %s", err)
		return err
	}

	var transactions []Transaction
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}
		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			log.Fatal(err)
			return err
		}
		layout := "01/01/2006"

		// Parse the string into a time.Time object
		parsedDate, err := time.Parse(layout, record[0])
		// Date Description Check# Category Currency Amount
		transaction := Transaction{
			Date:        parsedDate,
			Description: record[1],
			Category:    record[3],
			Amount:      price,
		}
		transactions = append(transactions, transaction)
	}

	err = a.InsertTransactions(transactions)
	return err
}

func (a *App) ParseTransactionsRange(after string, before string) ([]Transaction, error) {
	parsedBeforeDate, err := time.Parse("1/01/2006", before)
	if err != nil {
		return nil, err
	}
	// if there is no after date the front end will send 01/01/1900 to be safe
	parsedAfterDate, err := time.Parse("1/01/2006", after)
	if err != nil {
		return nil, err
	}
	transactions, err := a.GetTransactions() // db call
	for index, transaction := range transactions {
		if transaction.Date.After(parsedBeforeDate) || transaction.Date.Before(parsedAfterDate) {
			transactions = slices.Delete(transactions, index, index+1)
		}
	}

	return transactions, nil

}

func (a *App) GetCategoryTotals() (map[string]float64, error) {
	totalSpending = 0
	ignoreCategory := map[string]bool{
		"Buy":                 true,
		"Credit Card Payment": true,
		"Federal Tax":         true,
		"Fees & Charges":      true,
		"Transfer":            true,
		"Income":              true,
	}
	transactions, err := a.GetTransactions()
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}

	groupInfo := make(map[string]float64)
	for _, transaction := range transactions {
		if ignoreCategory[transaction.Category] {
			continue
		}
		groupInfo[transaction.Category] += transaction.Amount
		totalSpending += transaction.Amount
	}

	return groupInfo, nil
}

func (a *App) GetTotalSpending() float64 {
	return totalSpending
}
