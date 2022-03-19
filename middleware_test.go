package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticatedHTTPStatusCode(t *testing.T) {
	testCases := []struct {
		desc       string
		headers    map[string]string
		users      map[string]struct{}
		statusCode int
	}{
		{
			desc:       "Missing X-User-Id header",
			headers:    make(map[string]string),
			statusCode: 401,
		},
		{
			desc: "Empty X-User-Id header",
			headers: map[string]string{
				"X-User-Id": "",
			},
			statusCode: 401,
		},
		{
			desc: "Unknown user ID",
			headers: map[string]string{
				"X-User-Id": "unknown",
			},
			users: map[string]struct{}{
				"foo": {},
			},
			statusCode: 401,
		},
		{
			desc: "Known user ID",
			headers: map[string]string{
				"X-User-Id": "foo",
			},
			users: map[string]struct{}{
				"foo": {},
			},
			statusCode: 200,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// create an httprouter which uses the Authenticated middleware
			r := httprouter.New()
			r.GET("/", Authenticated(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				http.Error(w, "OK", 200)
			}, tc.users))

			// create a request with the given headers
			req := httptest.NewRequest("GET", "/", nil)
			for k, v := range tc.headers {
				req.Header.Set(k, v)
			}

			// execute the request and record the result
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			res := w.Result()

			assert.Equal(t, tc.statusCode, res.StatusCode)
		})
	}
}

func TestAuthenticatedSetsContextValue(t *testing.T) {
	expectedUserID := "user1"
	users := map[string]struct{}{
		expectedUserID: {},
	}

	// create an httprouter which uses the Authenticated middleware
	r := httprouter.New()
	r.GET("/", Authenticated(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		actualUserID := r.Context().Value(UserIDKey)
		assert.Equal(t, expectedUserID, actualUserID)
	}, users))

	// create a request with the given headers
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-User-Id", expectedUserID)

	// execute the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}
