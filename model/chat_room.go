// Model contains all the model for our Postgres database, just like ActiveRecord. We create a model and the snake case
// version of the struct name is going to be used as table name. Similarly, the snake cased field names will be used
// column names for the table.
package model

import "github.com/jinzhu/gorm"

// Default table name should be chat_rooms
type ChatRoom struct {
	gorm.Model
	Name     string
	Users    []User    `gorm:"many2many:memberships;"` // many-to-many
	Messages []Message `gorm:"ForeignKey:ChatRoomID"`  // has-many
}
