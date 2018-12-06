package main

import (
		"fmt"
	"time"
	"math"
)

func main() {
	moment := time.Now()

	fmt.Printf("The number: %v", moment.UnixNano())
	fmt.Println()
	fmt.Printf("The round: %v", int64(math.Round(float64(moment.UnixNano())/1000))*1000)
}