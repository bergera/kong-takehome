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

type getServiceResponse struct {
	ServiceID    string    `json:"serviceId"`
	OrgID        string    `json:"orgId"`
	Title        string    `json:"title"`
	Summary      string    `json:"summary"`
	VersionCount int       `json:"versionCount"`
	Versions     []Version `json:"versions"`
}

type getServiceVersionsResponse struct {
	ServiceID string    `json:"serviceId"`
	Count     int       `json:"count"`
	Versions  []Version `json:"versions"`
}

type getVersionResponse struct {
	ServiceID string `json:"serviceId"`
	VersionID string `json:"versionId"`
	Summary   string `json:"summary"`
}

func (ws *WebServer) GetServices(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := r.Context()

	services, err := ws.data.FindServices(ctx)
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

func (ws *WebServer) GetService(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := r.Context()
	serviceID := p.ByName("serviceID")

	service, err := ws.data.FindServiceByID(ctx, serviceID)
	if err != nil {
		fmt.Println("query failed: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if service == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	versions, err := ws.data.FindVersionsForService(ctx, serviceID)
	if err != nil {
		fmt.Println("query failed: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := getServiceResponse{
		ServiceID:    service.ServiceID,
		OrgID:        service.OrgID,
		Title:        service.Title,
		Summary:      service.Summary,
		VersionCount: len(versions),
		Versions:     versions,
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

func (ws *WebServer) GetServiceVersions(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := r.Context()
	serviceID := p.ByName("serviceID")

	service, err := ws.data.FindServiceByID(ctx, serviceID)
	if err != nil {
		fmt.Println("query failed: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if service == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	versions, err := ws.data.FindVersionsForService(ctx, serviceID)
	if err != nil {
		fmt.Println("query failed: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := getServiceVersionsResponse{
		ServiceID: serviceID,
		Count:     len(versions),
		Versions:  versions,
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

func (ws *WebServer) GetVersion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := r.Context()
	serviceID := p.ByName("serviceID")
	versionID := p.ByName("versionID")

	version, err := ws.data.FindVersionByID(ctx, serviceID, versionID)
	if err != nil {
		fmt.Println("query failed: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if version == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	resp := getVersionResponse{
		ServiceID: version.ServiceID,
		VersionID: version.VersionID,
		Summary:   version.Summary,
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
