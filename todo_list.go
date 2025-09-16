package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type task struct {
	id          int
	title       string
	description string
	isDone      bool
	category    string
	createdAt   time.Time
	expiredAt   time.Time
}

func openDB(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal("Error while opening the datatbase: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error while connecting to the datatbase: ", err)
	}

	fmt.Println("Connected!")
	return db, err
}
func creatingTable(name string, db *sql.DB) error {

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT UNIQUE,
		description TEXT,
		category TEXT,
		isDone INTEGER,
		createdAt TEXT,
		expiredAt TEXT
		);`)
	if err != nil {
		log.Fatal("Error while creating the table: ", err)
	}

	fmt.Println("The table is created")
	return err

}

func addTasks(tasks []task, db *sql.DB) {
	q := `INSERT INTO tasks (title, description, isDone, category, createdAt, expiredAt)
	      SELECT ?, ?, ?, ?, ?, ?
	      WHERE NOT EXISTS (SELECT 1 FROM tasks WHERE title = ?);`

	for _, t := range tasks {
		go func(t task) {
			_, err := db.Exec(q, t.title, t.description, t.isDone, t.category, t.createdAt, t.expiredAt, t.title)
			if err != nil {
				log.Println("Error inserting task:", t.title, err)
				// return err
			}
			fmt.Println("Inserted task:", t.title)
		}(t)

	}
	// return err
}
func removeTask(t task, title string, db *sql.DB) error {
	q := `DELETE FROM tasks WHERE title=$1`
	_, err := db.Exec(q, t.title)
	fmt.Println("The value has been removed successfully!")
	return err
}

func addToCat(t task, newCat string, db *sql.DB) error {
	q := `UPDATE tasks SET category=&1 WHERE title=$2`
	_, err := db.Exec(q, newCat, t.title)
	fmt.Println("The category has been update successfully!")
	return err
}

func markAsDone(t task, title string, db *sql.DB) error {
	q := `UPDATE tasks SET category='Done' WHERE title=$1`
	_, err := db.Exec(q, t.title)
	fmt.Println("The task has been marked as done successfully!")
	return err
}

func listByCat(cat string, db *sql.DB) {
	q := `SELECT title FROM tasks WHERE category=?`
	rows, err := db.Query(q, cat)
	if err != nil {
		log.Fatal("Error executing query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			log.Fatal("Error scanning row: ", err)
		}
		fmt.Println("Title: ", title)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Error with rows: ", err)
	}
}

// TODO
//func removeTable
//func removeDB

func main() {

	db, _ := openDB("todolist.db")
	creatingTable("tasks", db)

	task1 := task{
		title:       "task1",
		description: "description for task1",
		category:    "todo",
		isDone:      false,
		createdAt:   time.Now(),
		expiredAt:   time.Now().Add(time.Duration(time.Now().Day())),
	}

	task2 := task{
		title:       "task2",
		description: "description for task2",
		category:    "Done",
		isDone:      true,
		createdAt:   time.Now(),
		expiredAt:   time.Now().Add(time.Duration(time.Now().Day())),
	}

	task3 := task{
		title:       "task3",
		description: "description for task3",
		category:    "Done",
		isDone:      true,
	}

	task4 := task{
		title:       "task4",
		description: "description for task4",
		category:    "todo",
		isDone:      false,
	}

	task5 := task{
		title:       "task4",
		description: "description for task4",
		category:    "todo",
		isDone:      false,
	}

	addTasks([]task{task1, task2, task3, task4, task5}, db)
	time.Sleep(time.Second * 3)

	listByCat("todo", db)

	// err := removeTask(task1,"task1",db)
	// if err != nil {
	// 	log.Fatal("Error while removing an existing value: ", err)
	// }

	defer db.Close()

}
