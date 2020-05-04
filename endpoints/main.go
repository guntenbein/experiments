package main

import (
	"fmt"
	"sort"

	"golang.org/x/net/context"
)

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

type Middleware func(Endpoint) Endpoint

type PrioritisedMiddleware struct {
	Middleware
	priority int
}

type EndpointBuilder struct {
	Endpoint
	middlewares []PrioritisedMiddleware
}

func (eb *EndpointBuilder) AddMiddleware(pm PrioritisedMiddleware) {
	eb.middlewares = append(eb.middlewares, pm)
}

func (eb *EndpointBuilder) Build() Endpoint {
	sort.SliceStable(eb.middlewares, func(i, j int) bool {
		return eb.middlewares[i].priority < eb.middlewares[j].priority
	})
	e := eb.Endpoint
	for _, mw := range eb.middlewares {
		e = mw.Middleware(e)
	}
	return e
}

func main() {
	fmt.Printf("%+v", 10)
}
