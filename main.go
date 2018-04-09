package main

import (
	"github.com/ponyhoff/golocator/locator"
	"github.com/ponyhoff/golocator/rest"
	"log"
	"net/http"
	"os"
)

const (
	locationsDataFile = "data/locations.csv"
	networksDataFile  = "data/networks.csv"
)

func main() {
	handleError(
		"running locator service ...",
		run(),
	)
}

func handleError(msg string, err error) {
	if err != nil {
		log.Fatal(msg)
		os.Exit(2)
	}
}

func run() error {

	var (
		l = locator.NewLocator()
	)

	handleError(
		"loading data from csv...",
		locator.Load(
			locationsDataFile,
			networksDataFile,
			&l,
		),
	)

	ls := LocatorService{}

	mux := http.NewServeMux()
	mux.HandleFunc("/ip/:addr", rest.NewRESTHandler(
		parseURLParams(),
		getLocation(ls),
	))

	return http.ListenAndServe(":8080", mux)
}
