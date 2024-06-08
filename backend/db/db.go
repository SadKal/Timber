package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB{
	dsn := os.Getenv("DB_DSN")
	maxRetries := 10
    retryInterval := 5 * time.Second

	var db *gorm.DB
	var err error
	for retries := 0; retries < maxRetries; retries++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Print(err)
			time.Sleep(retryInterval)
			continue
		}else{
			log.Println("CONNECTED TO DB SUCCESSFULLY")
		}
        break
    }
	createTables(db)

	return db;
}