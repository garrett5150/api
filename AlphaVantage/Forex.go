package AlphaVantage

import "net/http"

//This API returns the realtime exchange rate for any pair of digital currency (e.g., Bitcoin) and physical currency (e.g., USD). Data returned for physical currency (Forex) pairs also include realtime bid and ask prices.
//Required Parameters
//function=CURRENCY_EXCHANGE_RATE
//from_currency=USD
//to_currency=CAD
//apikey=key
func FXExchangeRates(w http.ResponseWriter, r *http.Request) {

}

//This API returns intraday time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime.
//Required Parameters
//function=FX_INTRADAY
//from_symbol=CAD
//to_symbol=USD
//apikey=key
//interval=1min
//Optional Parameters
//datatype=json, csv
//outputsize=compact
func FXIntraDay(w http.ResponseWriter, r *http.Request) {

}

//This API returns the daily time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime.
//Required Parameters
//function=FX_DAILY
//from_symbol=CAD
//to_symbol=USD
//apikey=key
//interval=1min
//Optional Parameters
//datatype=json, csv
//outputsize=compact
func FXDaily(w http.ResponseWriter, r *http.Request) {

}

//This API returns the weekly time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime. The latest data point is the price information for the week (or partial week) containing the current trading day, updated realtime.
//Required Parameters
//function=FX_WEEKLY
//from_symbol=CAD
//to_symbol=USD
//apikey=key
//interval=1min
//Optional Parameters
//datatype=json, csv
func FXWeekly(w http.ResponseWriter, r *http.Request) {

}

//This API returns the monthly time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime. The latest data point is the prices information for the month (or partial month) containing the current trading day, updated realtime.
//Required Parameters
//function=FX_WEEKLY
//from_symbol=CAD
//to_symbol=USD
//apikey=key
//interval=1min
//Optional Parameters
//datatype=json, csv
func FXMonthly(w http.ResponseWriter, r *http.Request) {

}
