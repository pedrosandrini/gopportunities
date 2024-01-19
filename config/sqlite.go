package config

import (
	"os"

	"github.com/pedrosandrini/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not fond, creating ...")
		// creating the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create database and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.ErrorF("SQLite opening error: %v", err)
		return nil, err
	}

	// Migrate db schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("SQLite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
