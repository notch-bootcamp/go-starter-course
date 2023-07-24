package main

import (
	"fmt"
	"go-starter-course/basics/structs/types"
)

const Home = "Home"

type (
	LocationRegistry map[string]types.Vector3D
)

func main() {
	var locations LocationRegistry
	locations = make(LocationRegistry)

	locations[Home] = types.Vector3D{"Notch", 45.8093453, 15.9571886, 124}
	fmt.Printf("%+v\n", locations)

	var otherLocations = map[string]types.Vector3D{
		"Zero":   {},
		"Google": {"Google Offices", 37.42202, -122.08408, 8},
	}
	fmt.Printf("%+v\n", otherLocations)

	homeLocation := locations[Home]
	otherLocations[Home] = homeLocation
	fmt.Printf("%+v\n", otherLocations)

	delete(locations, Home)
	fmt.Printf("%+v\n", locations)

	if maybeHome, ok := locations[Home]; ok {
		fmt.Printf("Home is in locations: %v\n", maybeHome)
	} else {
		fmt.Println("Home is not in locations")
	}
}
