package AlphaVantage

//https://www.alphavantage.co/documentation/

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type StockReq struct {
	Symbol     string `json:"symbol"`
	Interval   string `json:"interval"`
	OutputSize string `json:"outputsize"`
	DataType   string `json:"datatype"`
}

type ApiKey struct {
	ApiKey string `json:"apikey"`
}

//This API returns intraday time series (timestamp, open, high, low, close, volume) of the equity specified.
//Required Parameters
//function=TIME_SERIES_INTRADAY
//symbol=IBM
//intreval=1min, 5min, 15min, 30min, 60min
//apikey=key
//Optional Parameters
//outputsize=compact, full
//datatype=json, csv
func TimeSeriesIntraday(w http.ResponseWriter, r *http.Request) {
	log.Info("Time Series Intraday Called")
	//Stock := StockReq{}
	var Stock StockReq
	var key ApiKey

	//reads the JSON from the request and assigns it to the Stock Structure
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read Body. ", err)
		return
	}
	err = json.Unmarshal(body, &Stock)
	if err != nil {
		log.Error("Failed Load Data to Stock Struct. ", err)
		return
	}
	log.WithFields(log.Fields{"Info Received": &Stock}).Info()

	//grabs the Alpha Vantage API Key
	file, err := ioutil.ReadFile("./AlphaVantage/AlphaVantageApiKey.json")
	if err != nil {
		log.Error("Failed to read ApiKey File. ", err)
		return
	}
	err = json.Unmarshal(file, &key)
	if err != nil {
		log.Error("Failed Load API Key to Struct. ", err)
		return
	}

	//TODO create the query to send to Alpha Vantage
	//TODO send query to Alpha Vantage and parse results
	//TODO replace the returns with proper HTTP responses

}
