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
