package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.WithFields(logrus.Fields{
		"name": "hh",
		"age":  18,
	}).Info("info msg")

	requestLogger := logrus.WithFields(logrus.Fields{
		"user_id": 10010,
		"ip":      "192.168.32.15",
	})

	requestLogger.Info("info msg")
	requestLogger.Error("error msg")
}
