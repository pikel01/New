package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "hello world")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}