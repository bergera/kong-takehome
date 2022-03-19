package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// create a new router with named parameters
	r := httprouter.New()
	r.GET("/services", GetServices)
	r.GET("/services/:serviceID", GetServiceByID)
	r.GET("/services/:serviceID/versions", GetVersionsByServiceID)
	r.GET("/services/:serviceID/versions/:versionID", GetVersionByID)

	// run the server without TLS on localhost
	log.Fatal(http.ListenAndServe(":8080", r))
}
