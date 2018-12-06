package main

import "fmt"

func main() {
	sample()
}

func sample() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}