package openstack

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var SECRET string

var OS_USERNAME string
var OS_PASSWORD string
var OS_PROJECT_NAME string
var OS_USER_DOMAIN_NAME string
var OS_PROJECT_DOMAIN_NAME string
var OS_AUTH_URL string
var OS_IDENTITY_API_VERSION string

var OS_DB_USERNAME string
var OS_DB_PASSWORD string
var OS_DB_IP string

var Log *log.Logger

var OS_DB *sql.DB

var PROHIBITED []string
var stmtUpdateUserPasswordHash *sql.Stmt

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
	Log = log.New(os.Stdout, "[openstack]", log.Ldate|log.Ltime|log.Llongfile)

	SECRET = secret

	OS_USERNAME = os.Getenv("OS_USERNAME")
	OS_PASSWORD = os.Getenv("OS_PASSWORD")
	OS_PROJECT_NAME = os.Getenv("OS_PROJECT_NAME")
	OS_USER_DOMAIN_NAME = os.Getenv("OS_USER_DOMAIN_NAME")
	OS_PROJECT_DOMAIN_NAME = os.Getenv("OS_PROJECT_DOMAIN_NAME")
	OS_AUTH_URL = os.Getenv("OS_AUTH_URL")
	OS_IDENTITY_API_VERSION = os.Getenv("OS_IDENTITY_API_VERSION")

	OS_DB_USERNAME = os.Getenv("OS_DB_USERNAME")
	OS_DB_PASSWORD = os.Getenv("OS_DB_PASSWORD")
	OS_DB_IP = os.Getenv("OS_DB_IP")

	OS_DB, err := sql.Open("mysql", OS_DB_USERNAME+":"+OS_DB_PASSWORD+"@tcp("+OS_DB_IP+":3306)/keystone")
	if err != nil {
		Log.Fatal(err)
	}
	err = OS_DB.Ping()
	if err != nil {
		Log.Fatal(err)
	}
	stmtUpdateUserPasswordHash, err = OS_DB.Prepare("UPDATE password SET password_hash = ? WHERE local_user_id = (SELECT id FROM local_user WHERE name = ?)")
	if err != nil {
		Log.Fatal(err)
	}

	PROHIBITED = []string{"admin", "glance", "placement", "neutron", "nova", "cinder"}

	m.HandleFunc("/api/openstack/user", checkSecretMiddleware(apiOpenstackHandler))
}

func Quit() {
	OS_DB.Close()
}
