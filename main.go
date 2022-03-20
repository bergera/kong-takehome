package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// mockUsers is a list of the user IDs we'll support for this exercise. In practice,
// these would be inferred from the database. Using map[string]struct{} as a lightweight set.
var mockUsers = map[string]struct{}{"1": {}, "2": {}}

func main() {
	// create a new router with named parameters
	r := httprouter.New()
	r.GET("/services", Authenticated(NotImplemented, mockUsers))
	r.GET("/services/:serviceID", Authenticated(NotImplemented, mockUsers))
	r.GET("/services/:serviceID/versions", Authenticated(NotImplemented, mockUsers))
	r.GET("/services/:serviceID/versions/:versionID", Authenticated(NotImplemented, mockUsers))

	// run the server without TLS on localhost
	log.Fatal(http.ListenAndServe(":8080", r))
}
