package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	constant "go-starter-course/web/location"
	"go-starter-course/web/types"
	"net/http"
)

const ServerURL = "http://localhost:9090"

func processErrorResponse(resp *http.Response) string {
	var errorResponse map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
		fmt.Println("Error while decoding error message:", err)
	}
	return errorResponse["error"]
}

func CreateLocation(location string, position types.Position) error {
	postUrl := fmt.Sprintf("%s/location?%s=%s", ServerURL, constant.LocationQueryParam, location)

	jsonBytes, err := json.Marshal(position)
	if err != nil {
		return err
	}
	resp, err := http.Post(postUrl, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Location was created:", location)
		return nil
	} else {
		errorMessage := processErrorResponse(resp)
		fmt.Printf("Status: %d, Error: %s\n", resp.StatusCode, errorMessage)
		return errors.New(fmt.Sprintf("error creating location: %s", errorMessage))
	}
}

func FetchLocation(location string) (*types.Position, error) {
	getUrl := fmt.Sprintf("%s/location?%s=%s", ServerURL, constant.KnownLocationQueryParam, location)
	resp, err := http.Get(getUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var position types.Position
		if err := json.NewDecoder(resp.Body).Decode(&position); err != nil {
			return nil, err
		}
		return &position, nil
	} else {
		errorMessage := processErrorResponse(resp)
		fmt.Printf("Status: %d, Error: %s\n", resp.StatusCode, errorMessage)
		return nil, errors.New(fmt.Sprintf("error fetching location: %s", errorMessage))
	}
}

func main() {
	err := CreateLocation("home", types.Position{45.8, 15.95, 124})
	if err != nil {
		fmt.Println("Didn't create home:", err)
	}
	err = CreateLocation("zero", types.Position{})
	if err != nil {
		fmt.Println("Didn't create zero:", err)
	}
	homePosition, err := FetchLocation("home")
	if err != nil {
		fmt.Println("Couldn't fetch home:", err)
		return
	}
	fmt.Println("Home is:", homePosition)
	googlePosition, err := FetchLocation("google")
	if err != nil {
		fmt.Println("Couldn't fetch google:", err)
		return
	}
	fmt.Println("Google is:", googlePosition)
}
