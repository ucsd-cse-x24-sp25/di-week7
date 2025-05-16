package main

import (
	"database/sql"
	"errors"
	"fmt"
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
 _, err := os.Stat(dbPath)
 dbExists := !errors.Is(err, os.ErrNotExist)


 db, err := sql.Open("sqlite3", dbPath)
 if err != nil {
   log.Fatalf("failed to open database: %s", err)
 }


 if !dbExists {
   createTableCommand := `
     CREATE TABLE students (
       id TEXT PRIMARY KEY,
       major TEXT NOT NULL
     );`
   _, err := db.Exec(createTableCommand)
   if err != nil {
     db.Close()
     log.Fatalf("failed to create students table: %s", err)
   }
   fmt.Println("Database and table created successfully.")
 }


 return db
}


// Task 2: Add a student record
func addStudent(db *sql.DB, id string, major string) {
 _, err := db.Exec("INSERT INTO students (id, major) VALUES (?, ?)", id, major)
 if err != nil {
   log.Fatal("Failed to insert entry:", err)
 }
 fmt.Printf("Inserted student: %s\n", id)
}


// Task 3: Read all student records
func readStudents(db *sql.DB) {
 rows, err := db.Query("SELECT id, major FROM students")
 if err != nil {
   log.Fatal("Failed to read from database:", err)
 }
 defer rows.Close()


 fmt.Println("Student list:")
 for rows.Next() {
   var id string
   var major string
   if err := rows.Scan(&id, &major); err != nil {
     log.Fatal(err)
   }
   fmt.Printf("ID: %s, Major: %s\n", id, major)
 }
}

