package main

import (
	"log"
	"net/http"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Context().Value("context_id"))

		next.ServeHTTP(w, r)
	})
}
func main() {}
