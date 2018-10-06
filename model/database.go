package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var PGConn *gorm.DB

// SetupDatabase will perform database connection and auto migration on all gorm.Models
func SetupDatabase() error {
	db, err := gorm.Open("postgres", "user=cfeng password=cfeng dbname=goyak sslmode=disable")
	if err != nil {
		return err
	}

	db.AutoMigrate(&User{}, &Message{}, &ChatRoom{})

	PGConn = db
	return nil
}
