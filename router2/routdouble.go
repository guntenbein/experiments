package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	h := func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	}
	r1 := mux.NewRouter()
	r2 := mux.NewRouter()
	r2.HandleFunc("/books/{title}/page/{page}", h)
	r1.PathPrefix("/books/{title}/page/{page}").Handler(r2)
	http.ListenAndServe(":80", r1)
}
