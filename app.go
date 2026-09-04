package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
)

var dbCon *sql.DB
var err error

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

func (a *App) ParseCsv(content string) ([]Transaction, error) {
	reader := csv.NewReader(strings.NewReader(content))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse CSV: %s", err)
		return nil, err
	}

	var transactions []Transaction
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}
		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			log.Fatal(err)
			return nil, err
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

	return transactions, nil
}

func (a *App) ParseTransactionsRange(transactions []Transaction, after string, before string) ([]Transaction, error) {
	parsedBeforeDate, err := time.Parse("01/01/2006", before)
	if err != nil {
		return nil, err
	}
	// if there is no after date the front end will send 01/01/1900 to be safe
	parsedAfterDate, err := time.Parse("01/01/2006", after)
	if err != nil {
		return nil, err
	}

	for index, transaction := range transactions {
		if transaction.Date.After(parsedBeforeDate) || transaction.Date.Before(parsedAfterDate) {
			transactions = slices.Delete(transactions, index, index+1)
		}
	}

	return transactions, nil

}

func (a *App) GetCategoryInfo() (map[string][]Transaction, map[string]float64, error) {
	transactions, err := a.GetTransactions()
	if err != nil {
		log.Fatal(err)
	}
	transactionGroups := make(map[string][]Transaction)
	groupInfo := make(map[string]float64)
	for _, transaction := range transactions {
		transactionGroups[transaction.Category] = append(transactionGroups[transaction.Category], transaction)
		if _, ok := groupInfo[transaction.Category]; ok {
			groupInfo[transaction.Category] = groupInfo[transaction.Category] + transaction.Amount
		} else {
			groupInfo[transaction.Category] = transaction.Amount
		}
	}
	return transactionGroups, groupInfo, nil
}
