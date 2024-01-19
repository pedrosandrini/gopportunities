package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return errors.New("Fake err")
}

func GetLogger(p string) *Logger {
	// initialize logger
	logger := NewLogger(p)
	return logger
}
