package main

import (
	"fmt"
	"go-starter-course/web/location"
	"net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do you want some locations?")
}

func main() {
	http.HandleFunc("/", simpleHandler)
	http.HandleFunc("/location", location.Handler)
	http.ListenAndServe(":9090", nil)
}
