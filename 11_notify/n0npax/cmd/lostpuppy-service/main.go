package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	puppy "github.com/anz-bank/go-course/11_notify/n0npax/pkg/puppy"
)

var (
	logFatalf = log.Fatal
	parser    = parseArgs
)

func main() {
	port, err := parser(os.Args[1:])
	if err != nil {
		logFatalf(err)
	}
	logFatalf(puppy.LostPuppyBackend().Run(fmt.Sprintf(":%d", port)))
}

func parseArgs(args []string) (int, error) {
	var port int
	kingpin.Flag("port", "Port number").Short('p').Default("8182").IntVar(&port)
	_, err := kingpin.CommandLine.Parse(args)
	return port, err
}
