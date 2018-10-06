package main

import (
	"goyak/model"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	Addr = ":3000"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	err := model.SetupDatabase()

	if err != nil {
		logrus.Error(err)
		return
	}

	defer model.PGConn.Close()

	server := &http.Server{
		Handler:      LoadRoutes(),
		Addr:         Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Infof("HTTP server is listening and serving on port %v", Addr)
	logrus.Fatal(server.ListenAndServe())
}
