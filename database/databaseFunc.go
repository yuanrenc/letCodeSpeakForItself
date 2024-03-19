package database

import (
	"database/sql"
	"fmt"
	"log"
)

type DatabaseConfig struct {
	DbUser     string
	DBPassword string
	DbPort     string
	DbName     string
	DbHost     string
}

type Task struct {
	Id       int
	Title    string
	Priority string
}

func ConnectToDataBase(dbConfig *DatabaseConfig) *sql.DB {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbUser, dbConfig.DBPassword)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
		return nil
	}
	fmt.Println("Connected to database")
	return db
}

func CreateTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title text NOT NULL,
			priority text NOT NULL
			)
			`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create a new table:", err)
	}
}

func InsertData(db *sql.DB) {
	tasks := []Task{
		{Title: "get a job", Priority: "high"},
		{Title: "buy milk", Priority: "low"},
		{Title: "take a coffee", Priority: "medium"}}
	for _, task := range tasks {
		query := `INSERT INTO tasks (title, priority) VALUES ($1, $2)`
		_, err := db.Exec(query, task.Title, task.Priority)
		if err != nil {
			log.Fatal("Failed to insert data into table:", err)
		}
	}
	fmt.Println("data was inserted")
}

func GetTasks(db *sql.DB) []Task {
	var tasks []Task
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal("Failed to retrieve data from table:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Title, &task.Priority)
		if err != nil {
			log.Fatal("Failed to scan row:", err)
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("Error occurred while iterating over rows:", err)
		return nil
	}
	return tasks
}
