package handler

import (
	"encoding/json"
	"fmt"
	"goyak/model"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func NewUserListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []model.User

		// if err := db.Find(&users).Error; err != nil {
		// 	RenderError(w, err.Error(), 400)
		// 	return
		// }

		res := []UserResponse{}
		for _, user := range users {
			userResponse := UserResponse{
				Name:  user.Name,
				Email: user.Email,
			}

			res = append(res, userResponse)
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

func NewUserRetrieveHandler() http.HandlerFunc {
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

type UserResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	SessionToken string `json:"session_token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewUserCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			logrus.Error(err)
		}

		u := &model.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}

		err := u.Create()
		if err != nil {
			switch err.(type) {
			case *model.DatabaseError:
				RenderError(w, fmt.Sprintf("encountered database error: %s", err.Error()), 422)
			case *model.ValidationError:
				RenderError(w, fmt.Sprintf("encountered validation error: %s", err.Error()), 422)
			default:
				RenderError(w, err.Error(), 500)
			}

			return
		}

		res := UserResponse{
			Name:         u.Name,
			Email:        u.Email,
			SessionToken: u.SessionToken,
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
