package task

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID          int
	Title       string
	Description string
	IsDone      bool
	Category    string
	CreatedAt   time.Time
	ExpiredAt   time.Time
}

func AddTasks(ctx context.Context, Tasks []Task, db *sql.DB) {
	q := `INSERT INTO Tasks (Title, Description, IsDone, Category, CreatedAt, ExpiredAt)
	      SELECT ?, ?, ?, ?, ?, ?
	      WHERE NOT EXISTS (SELECT 1 FROM Tasks WHERE Title = ?);`

	for _, t := range Tasks {
		go func(t Task, ctx context.Context) {
			_, err := db.ExecContext(ctx, q, t.Title, t.Description, t.IsDone, t.Category, t.CreatedAt, t.ExpiredAt, t.Title)
			if err != nil {
				log.Println("Error inserting task:", t.Title, err)
				// return err
			}
			fmt.Println("Inserted task:", t.Title)
		}(t, ctx)

	}
	// return err
}
func RemoveTask(t Task, title string, db *sql.DB) error {
	q := `DELETE FROM Tasks WHERE Title=$1`
	_, err := db.Exec(q, title)
	fmt.Println("The value has been removed successfully!")
	return err
}

func AddToCat(t Task, newCat string, db *sql.DB) error {
	q := `UPDATE Tasks SET Category=&1 WHERE Title=$2`
	_, err := db.Exec(q, newCat, t.Title)
	fmt.Println("The Category has been update successfully!")
	return err
}

func MarkAsDone(t Task, title string, db *sql.DB) error {
	q := `UPDATE Tasks SET Category='Done' WHERE Title=$1`
	_, err := db.Exec(q, title)
	fmt.Println("The Task has been marked as done successfully!")
	return err
}

func ListByCat(cat string, db *sql.DB) {
	q := `SELECT Title FROM Tasks WHERE Category=?`
	rows, err := db.Query(q, cat)
	if err != nil {
		log.Fatal("Error executing query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var Title string
		if err := rows.Scan(&Title); err != nil {
			log.Fatal("Error scanning row: ", err)
		}
		fmt.Println("Title: ", Title)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Error with rows: ", err)
	}
}
