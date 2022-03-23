package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// mockUsers is a list of the user IDs we'll support for this exercise. In practice,
// these would be inferred from the database. Using map[string]struct{} as a lightweight set.
var mockUsers = map[string]struct{}{"1": {}, "2": {}, "3": {}, "4": {}}

func main() {
	// establish the database connection parameters
	dsn := fmt.Sprintf("host=db user=kong_takehome password=%s dbname=kong_takehome sslmode=disable", os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// create our web server object which will hold shared dependencies
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
