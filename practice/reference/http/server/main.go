package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "item %s\n", r.PathValue("id"))
	})

	request := httptest.NewRequest(http.MethodGet, "/items/42", nil)
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: %d", recorder.Code))
	}
	fmt.Print(recorder.Body.String())
}
