package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetServices(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func GetServiceByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func GetVersionsByServiceID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func GetVersionByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
