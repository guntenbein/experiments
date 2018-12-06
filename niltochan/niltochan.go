package main

import "fmt"

func main() {
	type st struct {
	}
	var s st
	var c chan st
	c<-s
	fmt.Println(<-c)
	fmt.Println("Yaar!!")
}
