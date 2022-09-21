package openstack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func apiOpenstackHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		userCreateHandler(w, r)
	case http.MethodPut:
		userUpdateHandler(w, r)
	case http.MethodDelete:
		userDeleteHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getStatusCode(status string) int {
	code, _ := strconv.Atoi(strings.Fields(status)[0])
	return code
}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	var p OS_User_Create
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	driverRes := UserCreate(p.Email, p.PassworHash)
	w.WriteHeader(getStatusCode(driverRes.status))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, driverRes.body)
}

func userUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var p OS_User_Update
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	driverRes := UserUpdate(p.OldEmail, p.NewEmail, p.PasswordHash)
	w.WriteHeader(getStatusCode(driverRes.status))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, driverRes.body)
}

func userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var p OS_User_Delete

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	driverRes := UserDelete(p.Email)
	w.WriteHeader(getStatusCode(driverRes.status))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, driverRes.body)
}
