package model

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	UserID     int    `gorm:"index"` // Foreign key (belongs to)
	ChatRoomID int    `gorm:"index"` // Foreign key (belongs to)
	Body       string `gorm:"type:text"`
}
