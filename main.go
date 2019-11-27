package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My Basic HTTP Server")
	})

	http.ListenAndServe(":80", nil)
}
