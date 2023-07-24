package main

import (
	"github.com/gorilla/mux"
	"go-starter-course/web/location"
	"go-starter-course/web/middleware"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This home of router implementation"))
}

func main() {
	router := mux.NewRouter()

	getLocationHandler := http.HandlerFunc(location.GetKnownLocation)
	getLocationHandler = middleware.Logging(getLocationHandler)

	router.HandleFunc("/location", getLocationHandler).Methods("GET")
	router.HandleFunc(
		"/location",
		middleware.Logging(location.CreateLocation)).
		Methods("POST")
	router.HandleFunc("/", homeHandler).Methods("GET")

	http.ListenAndServe(":9091", router)
}
