package main

import (
	"context"
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type mockDataService struct {
	findServicesForUser func(ctx context.Context, userID int) ([]Service, error)
}

func (mock *mockDataService) FindServicesForUser(ctx context.Context, userID int) ([]Service, error) {
	return mock.findServicesForUser(ctx, userID)
}

func TestNotImplemented(t *testing.T) {
	// create an httprouter which uses the Authenticated middleware
	r := httprouter.New()
	r.GET("/", NotImplemented)

	// create a request with the given headers
	req := httptest.NewRequest("GET", "/", nil)

	// execute the request and record the result
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := w.Result()

	assert.Equal(t, 501, res.StatusCode)
}

func TestGetServices(t *testing.T) {
	testCases := []struct {
		desc        string
		ctx         context.Context
		data        DataService
		statusCode  int
		body        []byte
		contentType *string
	}{
		{
			desc:       "User ID not found in context",
			ctx:        context.TODO(),
			statusCode: 500,
		},
		{
			desc:       "User ID in context has incorrect data type",
			ctx:        context.WithValue(context.TODO(), UserIDKey, "foo"),
			statusCode: 500,
		},
		{
			desc: "FindServicesForUser returns error",
			ctx:  context.WithValue(context.TODO(), UserIDKey, 1),
			data: &mockDataService{
				findServicesForUser: func(ctx context.Context, userID int) ([]Service, error) {
					return nil, errors.New("data error")
				},
			},
			statusCode: 500,
		},
		{
			desc: "OK with empty result set",
			ctx:  context.WithValue(context.TODO(), UserIDKey, 1),
			data: &mockDataService{
				findServicesForUser: func(ctx context.Context, userID int) ([]Service, error) {
					return []Service{}, nil
				},
			},
			statusCode: 200,
			body:       []byte(`{"count":0,"services":[]}`),
		},
		{
			desc: "OK with results",
			ctx:  context.WithValue(context.TODO(), UserIDKey, 1),
			data: &mockDataService{
				findServicesForUser: func(ctx context.Context, userID int) ([]Service, error) {
					return []Service{
						{1, "Title 1", "Summary 1", 1, 1},
						{2, "Title 2", "Summary 2", 1, 1},
						{3, "Title 3", "Summary 3", 1, 1},
					}, nil
				},
			},
			statusCode: 200,
			body:       []byte(`{"count":3,"services":[{"serviceId":1,"title":"Title 1","summary":"Summary 1","orgId":1,"versionCount":1},{"serviceId":2,"title":"Title 2","summary":"Summary 2","orgId":1,"versionCount":1},{"serviceId":3,"title":"Title 3","summary":"Summary 3","orgId":1,"versionCount":1}]}`),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ws := &WebServer{
				data: tc.data,
			}

			// create an httprouter which uses the Authenticated middleware
			r := httprouter.New()
			r.GET("/", ws.GetServices)

			// create a request with the given headers
			req := httptest.NewRequest("GET", "/", nil)
			req = req.WithContext(tc.ctx)

			// execute the request and record the result
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			res := w.Result()

			assert.Equal(t, tc.statusCode, res.StatusCode)
			if tc.contentType != nil {
				assert.Equal(t, *tc.contentType, res.Header.Get("content-type"))
			}
			if tc.body != nil {
				body, _ := io.ReadAll(res.Body)
				t.Log(string(body))
				res.Body.Close()
				assert.Equal(t, tc.body, body)
			}
		})
	}
}
