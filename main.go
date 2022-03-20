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
	r.GET("/services", Recovery(Authenticated(NotImplemented, mockUsers)))
	r.GET("/services/:serviceID", Recovery(Authenticated(NotImplemented, mockUsers)))
	r.GET("/services/:serviceID/versions", Recovery(Authenticated(NotImplemented, mockUsers)))
	r.GET("/services/:serviceID/versions/:versionID", Recovery(Authenticated(NotImplemented, mockUsers)))

	// run the server without TLS on localhost
	log.Println("launching server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
