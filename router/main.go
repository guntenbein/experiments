package main

type Router struct{}

type SuperRouter struct {
	*Router
}

func main() {
	sr := SuperRouter{&Router{}}
	AddHTTPHandlers(sr.Router)
}

func AddHTTPHandlers(r *Router) {
}
