package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "students.db"

func main() {
	db := openDB()
	defer db.Close()

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run sqlite.go [add <id> <major> | read]")
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 4 {
			log.Fatal("Usage: add <id> <major>")
		}
		addStudent(db, os.Args[2], os.Args[3])
	case "read":
		readStudents(db)
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}

// Task 1: Open the DB or create it if it does not exist
func openDB() *sql.DB {
	panic("todo")
}

// Task 2: Add a student record
func addStudent(db *sql.DB, id string, major string) {
	panic("todo")
}

// Task 3: Read all student records
func readStudents(db *sql.DB) {
	panic("todo")
}