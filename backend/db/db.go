package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB{
	dsn := os.Getenv("DB_DSN")
	log.Println("DSN", dsn)
	maxRetries := 10
    retryInterval := 5 * time.Second

	var db *gorm.DB
	var err error
	for retries := 0; retries < maxRetries; retries++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Print(err)
			time.Sleep(retryInterval)
			continue
		}
        break // Connection successful, break out of loop
    }
	
	createTables(db)

	return db;
}