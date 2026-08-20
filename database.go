package main

import (
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

type TransactionRow struct {
	Date        time.Time
	Description string
	Category    string
	Amount      float64
}

func startDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./db.sqlite")
	if err != nil {
		return nil, err
	}

	// Verify the connection actually works
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	// Create a table matching TransactionRow
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS transactions (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		date        DATETIME NOT NULL,
		description TEXT NOT NULL,
		category    TEXT,
		amount      REAL NOT NULL
	)`)
	if err != nil {
		db.Close()
		return nil, err
	}

	log.Println("Connected and table ready")
	return db, nil
}

// InsertTransaction writes a single transaction to the db
func (a *App) InsertTransaction(t Transaction) error {
	_, err := dbCon.ExecContext(a.ctx,
		`INSERT INTO transactions (date, description, category, amount)
		 VALUES (?, ?, ?, ?)`,
		t.Date, t.Description, t.Category, t.Amount,
	)
	return err
}

// InsertTransactions writes a batch in a single transaction (fast + atomic)
func (a *App) InsertTransactions(transactions []Transaction) error {
	tx, err := dbCon.BeginTx(a.ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // no-op if committed

	stmt, err := tx.PrepareContext(a.ctx,
		`INSERT INTO transactions (date, description, category, amount)
		 VALUES (?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, t := range transactions {
		if _, err := stmt.ExecContext(a.ctx, t.Date, t.Description, t.Category, t.Amount); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// GetTransactions reads all rows back out
func (a *App) GetTransactions() ([]Transaction, error) {
	rows, err := dbCon.QueryContext(a.ctx,
		`SELECT date, description, category, amount FROM transactions ORDER BY date`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var t Transaction
		if err := rows.Scan(&t.Date, &t.Description, &t.Category, &t.Amount); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, rows.Err()
}
