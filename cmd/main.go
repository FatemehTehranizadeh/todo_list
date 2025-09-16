package main

import (
	"fmt"
	"log"
	"todo_list/internal/db"
	"todo_list/internal/task"
	"todo_list/pkg/utils"
)

func main() {

	utils.InitLogger()
	utils.InfoLogger.Println("Application started")

	database, err := db.OpenDB("todolist.db")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer database.Close()

	if err := db.CreateTable(database); err != nil {
		log.Fatal("Error creating table:", err)
	}

	t := task.Task{
		Title:       "First Task",
		Description: "This is the first task",
		Category:    "ToDo",
		IsDone:      false,
	}

	task.AddTasks([]task.Task{t}, database)
	// if err := task.AddTasks(t, database); err != nil {
	// 	log.Fatal("Error adding task:", err)
	// }

	fmt.Println("Task added successfully!")
}








// package main

// import (
// 	"context"
// 	"os"
// 	"time"

// 	"todo_list/internal/db"
// 	"todo_list/internal/task"
// 	"todo_list/pkg/utils"
// )

// func main() {
// 	// لاگر رو مقداردهی می‌کنیم
// 	if err := utils.InitLogger(); err != nil {
// 		// اگر لاگر نساخت نمیشه ادامه داد
// 		panic("failed to initialize logger: " + err.Error())
// 	}
// 	defer func() {
// 		_ = utils.Logger.Sync()
// 	}()

// 	utils.Logger.Info("Application started")

// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	// باز کردن دیتابیس
// 	database, err := db.OpenDB(ctx, "todolist.db")
// 	if err != nil {
// 		utils.Logger.Fatalf("Database connection failed: %v", err)
// 	}
// 	defer database.Close()

// 	// ساخت جدول
// 	if err := db.CreateTable(ctx, database); err != nil {
// 		utils.Logger.Fatalf("Error creating table: %v", err)
// 	}

// 	// چند تسک نمونه
// 	tasks := []task.Task{
// 		{
// 			Title:       "First Task",
// 			Description: "This is the first task",
// 			Category:    "ToDo",
// 			IsDone:      false,
// 			CreatedAt:   time.Now(),
// 			ExpiredAt:   time.Now().Add(24 * time.Hour),
// 		},
// 		{
// 			Title:       "Second Task",
// 			Description: "Another task example",
// 			Category:    "ToDo",
// 			IsDone:      false,
// 			CreatedAt:   time.Now(),
// 			ExpiredAt:   time.Now().Add(48 * time.Hour),
// 		},
// 	}

// 	// اضافه کردن با concurrency
// 	if err := task.AddTasks(ctx, tasks, database); err != nil {
// 		utils.Logger.Errorf("Error adding tasks: %v", err)
// 		os.Exit(1)
// 	}

// 	utils.Logger.Info("Tasks added successfully!")
// }