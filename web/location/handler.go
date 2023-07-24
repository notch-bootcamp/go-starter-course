package location

import (
	"encoding/json"
	"fmt"
	"go-starter-course/web/types"
	"net/http"
)

const (
	KnownLocationQueryParam = "knownLocation"
	LocationQueryParam      = "location"
)

var locationMap = make(map[string]types.Position)

func serializeError(errorMessage string) string {
	errorMap := map[string]string{"error": errorMessage}
	byteError, err := json.Marshal(errorMap)
	if err != nil {
		return "Internal error. Contact your administrator."
	}
	return string(byteError)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	http.Error(w, serializeError(message), status)
}

func GetKnownLocation(w http.ResponseWriter, r *http.Request) {
	knownLocation := r.URL.Query().Get(KnownLocationQueryParam)
	if len(knownLocation) == 0 {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("Missing %s parameter", KnownLocationQueryParam))
		return
	}
	var response interface{}
	position, ok := locationMap[knownLocation]
	if ok {
		response = position
		w.WriteHeader(http.StatusOK)
	} else {
		respondWithError(w, http.StatusBadRequest, "Unknown location")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func CreateLocation(w http.ResponseWriter, r *http.Request) {
	var location types.Position
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	locationLabel := r.URL.Query().Get(LocationQueryParam)
	if len(locationLabel) == 0 {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Missing %s parameter", LocationQueryParam))
		return
	}
	locationMap[locationLabel] = location
	w.WriteHeader(http.StatusCreated)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetKnownLocation(w, r)
	} else if r.Method == http.MethodPost {
		CreateLocation(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
