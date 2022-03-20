package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ContextKey struct{}

var UserIDKey ContextKey

const userIDHeader = "x-user-id"

// Authenticated is a middleware function that ensures the call is authenticated. If the
// call is authenticated, the user ID is set in the context. If not authenticated, responds
// with 401 Unauthorized.
func Authenticated(next httprouter.Handle, users map[string]struct{}) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// if a userID was provided in the X-UserId header and that userID
		// is in our list of mock users, then the request is considered authenticated.
		// In practice, we'd expect an auth token an the Authorization header and we
		// would validate its claims.
		userID := r.Header.Get(userIDHeader)
		if _, ok := users[userID]; ok {
			// set the user ID in the context so further handlers can access it
			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			r = r.WithContext(ctx)

			// call the next handler in the middleware chain
			next(w, r, p)
		} else {
			// if user ID not found, then respond with 401 Unauthorized
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

// Recovery responds with 500 Internal Server Error in the event of a panic.
func Recovery(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered from panic: ", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next(w, r, p)
	}
}
