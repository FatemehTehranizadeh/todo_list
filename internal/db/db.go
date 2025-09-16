package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"todo_list/pkg/utils"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(ctx context.Context, dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		utils.Logger.Errorf("error while opening the database:", err)
		return nil, fmt.Errorf("error while opening the database: %v", err)
	}

	ctxPing, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err = db.PingContext(ctxPing)
	if err != nil {
		utils.Logger.Errorf("error while connecting to the database: %v", err)
		db.Close()
		return nil, fmt.Errorf("error while connecting to the database: %v", err)
	}
	utils.Logger.Info("Connected to the database!")
	return db, nil
}

func CreateTable(ctx context.Context, db *sql.DB) error {
	q := `CREATE TABLE IF NOT EXISTS Tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		category TEXT,
		isDone INTEGER,
		createdAt TEXT,
		expiredAt TEXT
	);`
	_, err := db.ExecContext(ctx, q)
	if err != nil {
		utils.Logger.Errorf("error while creating the table: %v", err)
		db.Close()
		return fmt.Errorf("error while creating the table: %v", err)
	}
	utils.Logger.Info("Table has been created successfully!")
	return nil
}

// TODO
//func removeTable
//func removeDB
