package main

import (
	"net/http"

	"github.com/anz-bank/go-course/11_notify/willshen8/pkg/puppy"
	"github.com/sirupsen/logrus"
)

func main() {
	router := puppy.SetupRouter()
	puppy.SetupLostPuppyRoutes(router)
	logrus.Fatal(http.ListenAndServe(":8888", router))
}
