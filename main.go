package main

import (
	"github.com/sirupsen/logrus"
	"goyak/model"
)

func main() {
	if db, err := SetupDatabase(); err == nil {
		logrus.Infof("Perform checking if users table exists: %v", db.HasTable(&model.User{}))
		logrus.Infof("Perform checking if messages table exists: %v", db.HasTable(&model.Message{}))
		logrus.Infof("Perform checking if chat_rooms table exists: %v", db.HasTable(&model.ChatRoom{}))

		defer db.Close()
	} else {
		logrus.Error(err)
	}
}
