package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// mockUsers is a list of the user IDs we'll support for this exercise. In practice,
// these would be inferred from the database. Using map[string]struct{} as a lightweight set.
var mockUsers = map[string]struct{}{"1": {}, "2": {}, "3": {}, "4": {}}

func main() {
	db, err := sql.Open("postgres", "host=db user=kong_takehome password=abc123 dbname=kong_takehome sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	ws := &WebServer{
		data: &SQLDataService{
			db: db,
		},
	}

	// create a new router with named parameters
	r := httprouter.New()
	r.GET("/services", Recovery(Authenticated(ws.GetServices, mockUsers)))
	r.GET("/services/:serviceID", Recovery(Authenticated(ws.GetService, mockUsers)))
	r.GET("/services/:serviceID/versions", Recovery(Authenticated(ws.GetServiceVersions, mockUsers)))
	r.GET("/services/:serviceID/versions/:versionID", Recovery(Authenticated(ws.GetVersion, mockUsers)))

	// run the server without TLS on localhost
	log.Println("launching server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
