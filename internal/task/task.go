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



// package task

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	"todo_list/pkg/utils"
// )

// // Task ساختار قابل استفاده بیرون از پکیج.
// type Task struct {
// 	ID          int
// 	Title       string
// 	Description string
// 	IsDone      bool
// 	Category    string
// 	CreatedAt   time.Time
// 	ExpiredAt   time.Time
// }

// // AddTasks به صورت موازی تسک‌ها رو اضافه می‌کنه، خطاها برمی‌گردن و به context هم احترام می‌ذاره.
// func AddTasks(ctx context.Context, tasks []Task, db *sql.DB) error {
// 	type result struct {
// 		title string
// 		err   error
// 	}
// 	resultCh := make(chan result)
// 	defer close(resultCh)

// 	q := `INSERT INTO tasks (title, description, isDone, category, createdAt, expiredAt)
// 	      SELECT ?, ?, ?, ?, ?, ?
// 	      WHERE NOT EXISTS (SELECT 1 FROM tasks WHERE title = ?);`

// 	for _, t := range tasks {
// 		t := t // capture
// 		go func() {
// 			select {
// 			case <-ctx.Done():
// 				resultCh <- result{title: t.Title, err: ctx.Err()}
// 				return
// 			default:
// 			}

// 			_, err := db.ExecContext(ctx, q,
// 				t.Title,
// 				t.Description,
// 				boolToInt(t.IsDone),
// 				t.Category,
// 				t.CreatedAt.Format(time.RFC3339),
// 				t.ExpiredAt.Format(time.RFC3339),
// 				t.Title,
// 			)
// 			if err != nil {
// 				utils.Logger.Errorf("inserting task '%s' failed: %v", t.Title, err)
// 				resultCh <- result{title: t.Title, err: err}
// 				return
// 			}
// 			utils.Logger.Infof("Inserted task: %s", t.Title)
// 			resultCh <- result{title: t.Title, err: nil}
// 		}()
// 	}

// 	var firstErr error
// 	for i := 0; i < len(tasks); i++ {
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		case res := <-resultCh:
// 			if res.err != nil && firstErr == nil {
// 				firstErr = fmt.Errorf("task '%s' failed: %w", res.title, res.err)
// 			}
// 		}
// 	}
// 	return firstErr
// }

// func RemoveTask(ctx context.Context, title string, db *sql.DB) error {
// 	q := `DELETE FROM tasks WHERE title = ?`
// 	_, err := db.ExecContext(ctx, q, title)
// 	if err != nil {
// 		utils.Logger.Errorf("removing task '%s' failed: %v", title, err)
// 		return err
// 	}
// 	utils.Logger.Infof("Task '%s' removed successfully", title)
// 	return nil
// }

// func AddToCategory(ctx context.Context, title, newCat string, db *sql.DB) error {
// 	q := `UPDATE tasks SET category = ? WHERE title = ?`
// 	_, err := db.ExecContext(ctx, q, newCat, title)
// 	if err != nil {
// 		utils.Logger.Errorf("updating category for '%s' failed: %v", title, err)
// 		return err
// 	}
// 	utils.Logger.Infof("Category for '%s' updated to '%s'", title, newCat)
// 	return nil
// }

// func MarkAsDone(ctx context.Context, title string, db *sql.DB) error {
// 	q := `UPDATE tasks SET category = 'Done', isDone = 1 WHERE title = ?`
// 	_, err := db.ExecContext(ctx, q, title)
// 	if err != nil {
// 		utils.Logger.Errorf("marking task '%s' as done failed: %v", title, err)
// 		return err
// 	}
// 	utils.Logger.Infof("Task '%s' marked as done", title)
// 	return nil
// }

// func ListByCategory(ctx context.Context, cat string, db *sql.DB) error {
// 	q := `SELECT title FROM tasks WHERE category = ?`
// 	rows, err := db.QueryContext(ctx, q, cat)
// 	if err != nil {
// 		utils.Logger.Errorf("listing tasks by category '%s' failed: %v", cat, err)
// 		return fmt.Errorf("executing query: %w", err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var title string
// 		if err := rows.Scan(&title); err != nil {
// 			utils.Logger.Errorf("scanning row failed: %v", err)
// 			return fmt.Errorf("scanning row: %w", err)
// 		}
// 		fmt.Println("Title:", title) // می‌تونی اینو هم با logger بزنی اگر خواستی
// 	}
// 	if err := rows.Err(); err != nil {
// 		utils.Logger.Errorf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows error: %w", err)
// 	}
// 	return nil
// }

// func boolToInt(b bool) int {
// 	if b {
// 		return 1
// 	}
// 	return 0
// }