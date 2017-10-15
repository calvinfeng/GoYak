// Model contains all the model for our Postgres database, just like ActiveRecord. We create a model and the snake case
// version of the struct name is going to be used as table name. Similarly, the snake cased field names will be used
// column names for the table.
package model

import "github.com/jinzhu/gorm"

// Default table name is messages
type Message struct {
	gorm.Model
	UserID     int    `gorm:"index"` // Foreign key (belongs to)
	ChatRoomID int    `gorm:"index"` // Foreign key (belongs to)
	Body       string `gorm:"type:text"`
}
