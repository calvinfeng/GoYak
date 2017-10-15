package handler

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
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