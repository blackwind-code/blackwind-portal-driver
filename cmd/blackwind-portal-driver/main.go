package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/blackwind-code/blackwind-portal-driver/internal/openstack"
	"github.com/blackwind-code/blackwind-portal-driver/internal/zerotier"
)

var SECRET string

func ping(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Auth-Token") == SECRET {
		w.WriteHeader(200)
		fmt.Fprintf(w, "pong")
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func main() {
	SECRET = os.Getenv("SECRET")

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)
	openstack.Init(mux, SECRET)
	defer openstack.Quit()
	zerotier.Init(mux, SECRET)

	http.ListenAndServe(":8888", mux)
}
