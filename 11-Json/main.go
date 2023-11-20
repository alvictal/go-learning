package main

import (
	"encoding/json"
	"fmt"
)

// There is native resources to serialize and desearialize json objects, just follow the example below

type Car struct {
	Name  string
	Year  int
	Color string
}

func main() {

	car := Car{"BYD", 2023, "Cyan"}

	//Serializing the object
	result, _ := json.Marshal(car)
	fmt.Println((string(result)))

	var car2 Car
	json2 := []byte(`{"Name":"BYD Dolphin","Year":2023,"Color":"Black"}`)

	json.Unmarshal(json2, &car2)
	fmt.Println(car2)
}
