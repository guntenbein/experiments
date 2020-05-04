package main

import "fmt"

func returnMany() (int, string, error) {
	return 1, "example", nil
}

func main() {
	i, s, err := returnMany()
	fmt.Printf("Returned %s %s %v", i, s, err)
}
