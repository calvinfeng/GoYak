package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	Addr = ":3000"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	db, err := SetupDatabase()

	if err != nil {
		logrus.Error(err)
		return
	} else {
		defer db.Close()
	}

	server := &http.Server{
		Handler:      LoadRoutes(db),
		Addr:         Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Infof("Server is listening and serving on port %v", Addr)
	logrus.Fatal(server.ListenAndServe())
}
