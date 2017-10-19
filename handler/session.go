package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewSessionCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := []byte("secret")
		h := hmac.New(sha256.New, key) // Returns a hash function
		// h.Write([]byte(time.Now().String())) Feel free to write a message with a hash function
		logrus.Print(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	}
}
