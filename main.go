// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/akharb1192/LearningAgain/db" // Import db package for DBInterface
	_ "github.com/go-sql-driver/mysql"       // Import MySQL driver
)

// Global DBInterface variable
var DB db.DBInterface

// InitDB initializes the database connection (real or mock).
func InitDB(dbInstance db.DBInterface) {
	DB = dbInstance
}

// InsertUser inserts a new user into the database and returns the last inserted ID.
func InsertUser(name, email string) (int64, error) {
	// SQL statement for inserting a user
	sqlStatement := "INSERT INTO users (name, email) VALUES (?, ?)"

	// Call Exec on the actual DB interface instance (real or mock)
	result, err := DB.Exec(sqlStatement, name, email)
	if err != nil {
		return 0, fmt.Errorf("exec failed: %w", err)
	}

	// Retrieve the last insert ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	return lastInsertID, nil
}

func main() {
	// Real database connection (replace with your DB credentials)
	realDB, err := sql.Open("mysql", "root:ankit@tcp(127.0.0.1:3306)/ankit")
	if err != nil {
		log.Fatal(err)
	}
	defer realDB.Close()

	// Wrap *sql.DB in the SQLDBWrapper to implement DBInterface
	realDBWrapper := &db.SQLDBWrapper{DB: realDB}

	// Initialize the DB with the real database wrapper
	InitDB(realDBWrapper)

	// Call InsertUser function to insert a new user
	newUserID, err := InsertUser("Jane Doe", "jane.doe@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted new user with ID: %d\n", newUserID)
}
