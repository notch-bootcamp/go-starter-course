package main

import (
	"encoding/json"
	"fmt"
	"go-starter-course/web/types"
)

func main() {
	home := types.Position{45.8093453, 15.9571886, 124}
	homeJson, err := json.Marshal(home)
	if err != nil {
		fmt.Printf("Cannot marshal our variable: %s", err)
		return
	}
	fmt.Println(string(homeJson))
	homeJsonData := []byte(`{"lat":45.8093453,"long":15.9571886,"elev":124}`)
	var deserializedHome types.Position
	err = json.Unmarshal(homeJsonData, &deserializedHome)
	if err != nil {
		fmt.Printf("Cannon unmarshal our json: %s", err)
		return
	}
	fmt.Println(deserializedHome)
}
