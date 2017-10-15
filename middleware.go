package main

import (
	"net/http"
	"github.com/sirupsen/logrus"
)

type HttpMiddleware func(http.Handler) http.Handler

func NewServerLoggingMiddleware() HttpMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logrus.Infof("%s %s %s %s", r.Proto, r.Method, r.URL, r.Host)
			next.ServeHTTP(w, r)
		})
	}
}