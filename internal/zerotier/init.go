package zerotier

import (
	"log"
	"net/http"
	"os"
)

var ZEROTIER_API_URL string
var ZEROTIER_TOKEN string
var ZEROTIER_NODE_ID string
var ZEROTIER_NETWORK_ID string

var Log *log.Logger

func Init(m *http.ServeMux) {
	Log = log.New(os.Stdout, "[zerotier]", log.Ldate|log.Ltime|log.Llongfile)

	ZEROTIER_API_URL = os.Getenv("ZEROTIER_API_URL")
	ZEROTIER_TOKEN = os.Getenv("ZEROTIER_TOKEN")
	ZEROTIER_NODE_ID = os.Getenv("ZEROTIER_NODE_ID")
	ZEROTIER_NETWORK_ID = os.Getenv("ZEROTIER_NETWORK_ID")
}
