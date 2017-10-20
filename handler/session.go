package handler

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"goyak/model"
	"net/http"
)

type SessionResponse struct {
	Email        string `json:"email"`
	SessionToken string `json:"session_token"`
}

func NewSessionCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, password := r.PostFormValue("email"), r.PostFormValue("password")
		if len(email) == 0 || len(password) == 0 {
			RenderError(w, "Please Provide email and password for authentication", 400)
			return
		}

		var user model.User
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			RenderError(w, "Email is not recognized", 400)
			return
		}

		if err := bcrypt.CompareHashAndPassword(user.PasswordDigest, []byte(password)); err != nil {
			RenderError(w, "Incorrect password", 400)
		} else {
			res := SessionResponse{
				Email:        user.Email,
				SessionToken: user.SessionToken,
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
}
