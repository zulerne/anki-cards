package main

import (
	"fmt"
	"net/http"
)

func main() {}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		next.ServeHTTP(w, r)
	})
}
