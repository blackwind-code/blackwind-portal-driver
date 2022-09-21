package zerotier

import (
	"log"
	"net/http"
	"os"
)

var SECRET string

var ZEROTIER_API_URL string
var ZEROTIER_TOKEN string
var ZEROTIER_NODE_ID string
var ZEROTIER_NETWORK_ID string

var Log *log.Logger

var PROHIBITED []string

func checkSecretMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Auth-Token") == SECRET {
			next(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func Init(m *http.ServeMux, secret string) {
	Log = log.New(os.Stdout, "[zerotier]", log.Ldate|log.Ltime|log.Llongfile)

	SECRET = secret

	ZEROTIER_API_URL = os.Getenv("ZEROTIER_API_URL")
	ZEROTIER_TOKEN = os.Getenv("ZEROTIER_TOKEN")
	ZEROTIER_NODE_ID = os.Getenv("ZEROTIER_NODE_ID")
	ZEROTIER_NETWORK_ID = os.Getenv("ZEROTIER_NETWORK_ID")

	PROHIBITED = []string{"04055d4bbe"}

	m.HandleFunc("/api/zerotier/device", checkSecretMiddleware(apiZerotierHandler))
}
