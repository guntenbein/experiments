package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type ss struct {
		A *int `json:"a"`
		B *int `json:"b"`
	}

	d := `{"a":11}`

	s := &ss{}

	fmt.Printf("%s\n",d)

	err := json.Unmarshal([]byte(d),s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n",s)

}
