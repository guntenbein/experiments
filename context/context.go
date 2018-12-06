package main

import (
	"context"
	"fmt"
)

type contextKey string

func main() {
	ctx := context.Background()
	ctxUp := context.WithValue(ctx,contextKey("a"),"gross")
	fmt.Println(ctxUp.Value(contextKey("a")))
}
