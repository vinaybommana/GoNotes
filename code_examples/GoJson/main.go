package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Message struct {
		Name string
		Body string
		Time int64
	}

	m := Message{"Alice", "Hello", 129}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error!")
	}

	fmt.Println(string(b))

	// Unmarshalling data
	// var unmarshalMessage Message
	// err := json.Unmarshal(b, &unmarshalMessage)
}
