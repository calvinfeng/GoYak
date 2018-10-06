// Model contains all the model for our Postgres database, just like ActiveRecord. We create a model and the snake case
// version of the struct name is going to be used as table name. Similarly, the snake cased field names will be used
// column names for the table.
package model

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name           string     `gorm:"type:varchar(100)" json:"name"`
	Email          string     `gorm:"type:varchar(100);unique_index" json:"email"` // Set SQL type and create unique indexing
	SessionToken   string     `gorm:"type:varchar(100);unique_index"`
	PasswordDigest []byte     `gorm:"type:bytea;unique_index"`
	ChatRooms      []ChatRoom `gorm:"many2many:memberships"` // many-to-many
	Messages       []Message  `gorm:"ForeignKey:UserID"`     // has-many
	Password       string     `gorm:"-"`
}

func (u *User) Create() error {
	if PGConn == nil {
		return fmt.Errorf("PostgreSQL is not connected")
	}

	if len(u.Email) < 5 {
		return NewValidationError("email", "email is too short")
	}

	if len(u.Password) < 5 {
		return NewValidationError("password", "password is too short")
	}

	var err error
	u.PasswordDigest, err = bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	u.SessionToken, err = GenerateRandomString(64)
	if err != nil {
		return err
	}

	err = PGConn.Create(u).Error
	if err != nil {
		switch {
		case err == gorm.ErrRecordNotFound:
			return &DatabaseError{fmt.Sprintf("record is not found %s", err.Error())}
		case err == gorm.ErrInvalidSQL:
			return &DatabaseError{fmt.Sprintf("bad sql %s", err.Error())}
		case err == gorm.ErrInvalidTransaction:
			return &DatabaseError{fmt.Sprintf("bad transaction %s", err.Error())}
		case err == gorm.ErrCantStartTransaction:
			return &DatabaseError{fmt.Sprintf("cannot start transactionl %s", err.Error())}
		case err == gorm.ErrUnaddressable:
			return &DatabaseError{fmt.Sprintf("unaddressable %s", err.Error())}
		default:
			return &DatabaseError{fmt.Sprintf("I dont fucking know %s", err.Error())}
		}
	}

	return nil
}

func GenerateRandomBytes(n int) ([]byte, error) {
	byteArray := make([]byte, n)
	_, err := rand.Read(byteArray)
	if err != nil {
		return nil, err
	}

	return byteArray, nil
}

func GenerateRandomString(length int) (string, error) {
	bytes, err := GenerateRandomBytes(length)
	return base64.URLEncoding.EncodeToString(bytes), err
}
