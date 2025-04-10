package database

import (
	"fmt"
	"log"

	models "github.com/AankTia/go-simple-crud/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connects to the SQLite databse
func ConnectDB()  {
	var err error

	// Open a connection to an SQLite database file named "tasks.db"
	// If file doesn'r exist, it will be created
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	fmt.Println("Connected to SQLite database successfully")

	// Auto Migrate the Task model to the database
	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migration completed")
}