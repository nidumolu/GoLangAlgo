package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"username"`
	Lastname string `json:"lastname"`
	Year     int    `json:"birth_year"`
}

func main() {
	useall := User{Name: "John", Lastname: "Adams", Year: 2000}
	// Regular Structure
	// Encoding JSON data -> Convert Go Structure to JSON record with fields
	t, err := json.Marshal(&useall)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Json Marshalling a Struct , Value %s\n", t)
	}
	// Decoding JSON data given as a string
	str := `{"username": "John", "lastname": "Doe", "birth_year":2001}`

	// Convert string into a byte slice
	jsonRecord := []byte(str)

	// Create a structure variable to store the result
	temp := User{}
	err = json.Unmarshal(jsonRecord, &temp)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Json Unmarshall : Data type: %T with value %v\n", temp, temp)
	}

}
