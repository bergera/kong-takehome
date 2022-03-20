package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type WebServer struct {
	data DataService
}

type getServicesResponse struct {
	Count    int       `json:"count"`
	Services []Service `json:"services"`
}

func (ws *WebServer) GetServices(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := r.Context()

	userID, ok := ctx.Value(UserIDKey).(int)
	if !ok {
		fmt.Println("user ID not found in context")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	services, err := ws.data.FindServicesForUser(ctx, userID)
	if err != nil {
		fmt.Println("query failed: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := getServicesResponse{
		Count:    len(services),
		Services: services,
	}
	body, err := json.Marshal(&resp)
	if err != nil {
		fmt.Println("marshal response body failed: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func NotImplemented(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
