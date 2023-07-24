package location

import (
	"bytes"
	"encoding/json"
	"go-starter-course/web/types"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetKnownLocation(t *testing.T) {
	locationMap["knownLocation"] = types.Position{
		Latitude:  12.345,
		Longitude: 67.89,
		Elevation: 100,
	}

	testCases := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Known location",
			queryParams:    "knownLocation=knownLocation",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"lat":12.345,"long":67.89,"elev":100}`,
		},
		{
			name:           "Missing parameter in URL Query",
			queryParams:    "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Missing knownLocation parameter"}`,
		},
		{
			name:           "Unknown location",
			queryParams:    "knownLocation=none",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Unknown location"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodGet, "/?"+tc.queryParams, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				GetKnownLocation(w, r)
			})
			handler.ServeHTTP(recorder, request)

			if recorder.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tc.expectedStatus, recorder.Code)
			}
			responseBody := strings.TrimSpace(recorder.Body.String())
			if responseBody != tc.expectedBody {
				t.Errorf("Expected response body %s, but got %s", tc.expectedBody, responseBody)
			}
		})
	}
}

func TestCreateLocation(t *testing.T) {
	testCases := []struct {
		name           string
		queryParams    string
		requestBody    types.Position
		expectedStatus int
	}{
		{
			name:           "Successful creation",
			queryParams:    "location=newLocation",
			requestBody:    types.Position{Latitude: 12.345, Longitude: 67.89, Elevation: 100},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "Missing parameter creation",
			queryParams:    "",
			requestBody:    types.Position{Latitude: 12.345, Longitude: 67.89, Elevation: 100},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			locationJson, err := json.Marshal(tc.requestBody)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}

			request, err := http.NewRequest(http.MethodPost, "/?"+tc.queryParams, bytes.NewBuffer(locationJson))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(CreateLocation)
			handler.ServeHTTP(recorder, request)

			if recorder.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tc.expectedStatus, recorder.Code)
			}
		})
	}
}
