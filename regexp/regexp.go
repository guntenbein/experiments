package main

import (
	"regexp"
	"fmt"
)

func main() {
	fmt.Println(regexp.MatchString("\\.go$", "git/logs/refs/remotes/origin/kafka_support_into_gokit"))
}