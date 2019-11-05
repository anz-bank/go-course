package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/anz-bank/go-course/11_notify/kasunfdo/pkg/puppy"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	port := kingpin.Flag("port", "Lost Puppy service port").Short('p').Default("8081").Int()
	_, err := kingpin.CommandLine.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	lostAPIHandler := puppy.NewLostAPIHandler()
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	lostAPIHandler.WireRoutes(router)

	serverAddr := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
