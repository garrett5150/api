package main

import (
	"api/AlphaVantage"
	"github.com/gorilla/mux"
)

//Routes Incoming requests to various functions
func routes(r *mux.Router) {
	//Stock Time Series Routes
	r.HandleFunc("/AlphaVantage/STS/TimeSeriesIntraday", AlphaVantage.TimeSeriesIntraday).Methods("POST")
	r.HandleFunc("/AlphaVantage/STS/TimeSeriesDaily", AlphaVantage.TimeSeriesDaily).Methods("POST")
	r.HandleFunc("/AlphaVantage/STS/TimeSeriesDailyAdjusted", AlphaVantage.TimeSeriesDailyAdjusted).Methods("POST")
	r.HandleFunc("/AlphaVantage/STS/TimeSeriesWeekly", AlphaVantage.TimeSeriesWeekly).Methods("POST")
	r.HandleFunc("/AlphaVantage/STS/TimeSeriesWeeklyAdjusted", AlphaVantage.TimeSeriesWeeklyAdjusted).Methods("POST")
	r.HandleFunc("/AlphaVantage/STS/TimeSeriesMonthly", AlphaVantage.TimeSeriesMonthly).Methods("POST")
	r.HandleFunc("/AlphaVantage/STS/TimeSeriesMonthlyAdjusted", AlphaVantage.TimeSeriesMonthlyAdjusted).Methods("POST")
	r.HandleFunc("/AlphaVantage/STS/QuoteEndpoint", AlphaVantage.QuoteEndpoint).Methods("POST")
	r.HandleFunc("/AlphaVantage/STS/SearchEndpoint", AlphaVantage.SearchEndpoint).Methods("POST")
	//Forex Routes
	r.HandleFunc("/AlphaVantage/FX/ExchangeRates", AlphaVantage.FXExchangeRates).Methods("POST")
	r.HandleFunc("/AlphaVantage/FX/IntraDay", AlphaVantage.FXIntraDay).Methods("POST")
	r.HandleFunc("/AlphaVantage/FX/Daily", AlphaVantage.FXDaily).Methods("POST")
	r.HandleFunc("/AlphaVantage/FX/Weekly", AlphaVantage.FXWeekly).Methods("POST")
	r.HandleFunc("/AlphaVantage/FX/Monthly", AlphaVantage.FXMonthly).Methods("POST")
	//Cyrpto Routes
	r.HandleFunc("/AlphaVantage/Crypto/ExchangeRates", AlphaVantage.CryptoExchangeRates).Methods("POST")
	r.HandleFunc("/AlphaVantage/Crypto/CryptoHealthIndex", AlphaVantage.CryptoHealthIndex).Methods("POST")
	r.HandleFunc("/AlphaVantage/Crypto/Daily", AlphaVantage.CryptoDaily).Methods("POST")
	r.HandleFunc("/AlphaVantage/Crypto/Weekly", AlphaVantage.CryptoWeekly).Methods("POST")
	r.HandleFunc("/AlphaVantage/Crypto/Monthly", AlphaVantage.CryptoMonthly).Methods("POST")
}
