package main

import (
	"api/AlphaVantage"
	"github.com/gorilla/mux"
)

//Routes Incoming requests to various functions
func routes(r *mux.Router) {
	r.HandleFunc("/", homeLink).Methods("GET")
	r.HandleFunc("/test1", test1)
	r.HandleFunc("/test2", test2).Methods("GET") //.Host("http://LocalHost:8080")
	r.HandleFunc("/AlphaVantage/STS/TimeSeriesIntraday", AlphaVantage.TimeSeriesIntraday).Methods("POST")
	r.HandleFunc("/AlphaVantage/STS/TimeSeriesIntraday", AlphaVantage.TimeSeriesIntraday).Methods("POST")
}
