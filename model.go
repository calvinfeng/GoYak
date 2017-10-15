package main

import "github.com/jinzhu/gorm"

/*
	Column name is the snake case of field's name
*/

// Default table name is users
type User struct {
	gorm.Model
	Name   string
	Emails []Email // has many
}

// Default table name is emails
type Email struct {
	gorm.Model
	UserID     int    `gorm:"index"`                          // Foreign key (belongs to)
	Email      string `gorm:"type:varchar(100);unique_index"` // Set SQL type and create unique indexing
	Subscribed bool
}

// Default table name is messages
type Message struct {
	gorm.Model
	UserID int    `gorm:"index"`
	Body   string `gorm:"type:text"`
}
