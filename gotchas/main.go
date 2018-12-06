package main

import "fmt"

func main() {
	fmt.Println(example2())
}

func example2() (i int) {
	defer func(){
		i = 3
	}()
	return 2
}
