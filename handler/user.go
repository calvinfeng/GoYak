package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"goyak/model"
	"net/http"
)

func NewUserListHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User list is here!"))
	}
}

func NewUserRetrieveHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var responseText string

		params := mux.Vars(r)
		if params["id"] != "" {
			responseText = fmt.Sprintf("User %v detail is here!", params["id"])
		} else {
			responseText = "User id is not provided."
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseText))
	}
}

// if err := bcrypt.CompareHashAndPassword(hashBytes, []byte("12367")); err != nil {
// 	logrus.Error("Incorrect password")
// } else {
// 	logrus.Print("Password is correct")
// }

func NewUserCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var responseByteArray []byte

		if err := r.ParseForm(); err != nil {
			logrus.Error(err)
		}

		name, email, password := r.PostFormValue("name"), r.PostFormValue("email"), r.PostFormValue("password")
		 if len(email) == 0 || len(password) == 0 {
		 	http.Error(w, "Please provide email and password for user sign up", 400)
		 	return
		 }

		if hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10); err == nil {
			newUser := model.User{
				Name:           name,
				Email:          email,
				PasswordDigest: hashBytes,
			}

			if err := db.Create(&newUser).Error; err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			if bytes, err := json.Marshal(newUser); err == nil {
				responseByteArray = bytes
			} else {
				http.Error(w, err.Error(), 500)
				return
			}
		} else {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseByteArray)
	}
}
