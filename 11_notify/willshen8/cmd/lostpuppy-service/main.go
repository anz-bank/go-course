package main

import (
	"net/http"
	"time"

	"github.com/anz-bank/go-course/11_notify/willshen8/pkg/puppy"
	"github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
)

var (
	// graceful library shuts down the server for testing purposes
	server = &graceful.Server{
		Timeout: 3 * time.Second,
	}
	stop     chan bool
	testMode = false
)

// runMain() stops the server after 3 seconds under test mode
func runMain() {
	// start the stop channel
	stop = make(chan bool)
	// put the service in "testMode"
	testMode = true
	// run the main entry point
	go main()
	// watch for the stop channel
	<-stop
	// stop the graceful server
	server.Stop(3 * time.Second)
}

func main() {
	if testMode {
		router := puppy.SetupRouter()
		puppy.SetupLostPuppyRoutes(router)
		stop <- true
		logrus.Fatal(http.ListenAndServe(":8888", router))
	}
}
