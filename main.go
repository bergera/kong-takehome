package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// mockUsers is a list of the user IDs we'll support for this exercise. In practice,
// these would be inferred from the database. Using map[string]struct{} as a lightweight set.
var mockUsers = map[string]struct{}{"foo": {}, "bar": {}}

func main() {
	// create a new router with named parameters
	r := httprouter.New()
	r.GET("/services", Authenticated(GetServices, mockUsers))
	r.GET("/services/:serviceID", Authenticated(GetServiceByID, mockUsers))
	r.GET("/services/:serviceID/versions", Authenticated(GetVersionsByServiceID, mockUsers))
	r.GET("/services/:serviceID/versions/:versionID", Authenticated(GetVersionByID, mockUsers))

	// run the server without TLS on localhost
	log.Fatal(http.ListenAndServe(":8080", r))
}
