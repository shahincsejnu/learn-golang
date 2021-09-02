package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func main() {
	fmt.Println("Welcome to json decoding")

	jsonString := `{
					"name" : "battery sensor",
					"capacity" : 40,
					"time" : "2020-12-21:12:52pm"
				}`

	var reading SensorReading

	err := json.Unmarshal([]byte(jsonString), &reading)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)
}
