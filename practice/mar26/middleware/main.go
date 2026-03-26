package main

import (
	"log"
	"net/http"
)

func main() {}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Context().Value("request_id"))

		next.ServeHTTP(w, r)
	})
}
