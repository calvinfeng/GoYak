package handler

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

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

func RenderError(w http.ResponseWriter, message string, code int) {
	res := ErrorResponse{
		Error: message,
	}

	bytes, _ := json.Marshal(res)
	w.WriteHeader(code)
	w.Write(bytes)
}
