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

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewUserCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			logrus.Error(err)
		}

		name, email, password := r.PostFormValue("name"), r.PostFormValue("email"), r.PostFormValue("password")
		if len(email) == 0 || len(password) == 0 {
			RenderError(w, "Please Provide email and password for user sign up", 400)
			return
		}

		hashBytes, hashErr := bcrypt.GenerateFromPassword([]byte(password), 10)
		if hashErr != nil {
			RenderError(w, hashErr.Error(), 500)
			return
		}

		newUser := model.User{
			Name:           name,
			Email:          email,
			PasswordDigest: hashBytes,
		}

		if err := db.Create(&newUser).Error; err != nil {
			RenderError(w, err.Error(), 400)
			return
		}

		res := UserResponse{
			Name:  newUser.Name,
			Email: newUser.Email,
		}

		if bytes, err := json.Marshal(res); err != nil {
			RenderError(w, err.Error(), 500)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}
	}
}

func RenderError(w http.ResponseWriter, message string, code int) {
	res := ErrorResponse{
		Error: message,
	}

	bytes, _ := json.Marshal(res)
	w.WriteHeader(code)
	w.Write(bytes)
}
