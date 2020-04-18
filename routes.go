package main

import "github.com/gorilla/mux"

//Routes Incoming requests to various functions
func routes(r *mux.Router) {
	r.HandleFunc("/", homeLink)
	r.HandleFunc("/test1", test1)
	r.HandleFunc("/test2", test2)
}
