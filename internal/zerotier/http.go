package zerotier

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func apiZerotierHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		deviceCreateHandler(w, r)
	case http.MethodPut:
		deviceUpdateHandler(w, r)
	case http.MethodDelete:
		deviceDeleteHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getStatusCode(status string) int {
	code, _ := strconv.Atoi(strings.Fields(status)[0])
	return code
}

func deviceCreateHandler(w http.ResponseWriter, r *http.Request) {
	var p Device_Create
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	driverRes := DeviceCreate(p.ZTAddress)
	w.WriteHeader(getStatusCode(driverRes.status))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, driverRes.body)
}

func deviceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var p Device_Update
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	driverRes := DeviceUpdate(p.ZTAddress, p.DeviceType)
	w.WriteHeader(getStatusCode(driverRes.status))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, driverRes.body)
}

func deviceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var p Device_Delete
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	driverRes := DeviceDelete(p.ZTAddress)
	w.WriteHeader(getStatusCode(driverRes.status))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, driverRes.body)
}
