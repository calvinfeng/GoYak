package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := gorm.Open(
		"postgres",
		"user=cfeng password=cfeng dbname=goyak_development sslmode=disable",
	)

	if err != nil {
		logrus.Error(err)
	} else {
		// Perform auto migrations
		db.AutoMigrate(&User{}, &Email{}, &Message{})
	}

	logrus.Infof("Perform checking if users table exists: %v", db.HasTable(&User{}))
	logrus.Infof("Perform checking if emails table exists: %v", db.HasTable(&Email{}))
	logrus.Infof("Perform checking if messages table exists: %v", db.HasTable(&Message{}))

	defer db.Close()
}
