package main

import (
	"log"
	"net/http"
)

func checkHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check security header
		next.ServeHTTP(w, r)
	})
}

func init() {

}

func main() {
	log.Println("Hello, World!")
}
