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
	OutputSize string `json:"outputsize,omitempty"`
	DataType   string `json:"datatype,omitempty"`
}

type ApiKey struct {
	ApiKey string `json:"apikey"`
}

type Error struct {
	Error string `json:"error"`
}

//TODO Change the strings to either some int var or a Money value
type IntraDayStock struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

//TODO This structure is not confirmed to be working yet
type IntraDay struct {
	MetaData struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		Interval      string `json:"4. Interval"`
		OutputSize    string `json:"5. Output Size"`
		TimeZone      string `json:"6. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries struct {
		Stock map[string]IntraDayStock
	} `json:"Time Series (15min)"`
}

var GenericError Error
var AlphaVantageURL = "https://www.alphavantage.co/query?"

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
	var returnValue IntraDay
	var Stock StockReq
	var key ApiKey

	//reads the JSON from the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read Body. ", err)
		//Set Content-type & Status to client can read the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		//Prep Error to be sent back to requester
		GenericError.Error = "The server could not understand the request due to invalid syntax"
		resErr, _ := json.Marshal(GenericError)
		//Write the response back to requester
		w.Write(resErr)
		return
	}
	//Assigns JSON Request to the Stock Structure
	err = json.Unmarshal(body, &Stock)
	if err != nil {
		log.Error("Failed Load Data to Stock Struct. ", err)
		//Set Content-type & Status to client can read the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		//Prep Error to be sent back to requester
		GenericError.Error = "The server could not understand the request due to invalid syntax"
		resErr, _ := json.Marshal(GenericError)
		//Write the response back to requester
		w.Write(resErr)
		return
	}
	//Log the info we received
	log.WithFields(log.Fields{"Info Received": &Stock}).Info()

	//grabs the Alpha Vantage API Key
	file, err := ioutil.ReadFile("./AlphaVantage/AlphaVantageApiKey.json")
	if err != nil {
		log.Error("Failed to read ApiKey File. ", err)

		//Set Content-type & Status to client can read the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		//Prep Error to be sent back to requester
		GenericError.Error = "The service is currently unavailable"
		resErr, _ := json.Marshal(GenericError)
		//Write the response back to requester
		w.Write(resErr)
		return
	}
	//Assigns the API key to the struct
	err = json.Unmarshal(file, &key)
	if err != nil {
		log.Error("Failed Load API Key to Struct. ", err)

		//Set Content-type & Status to client can read the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		//Prep Error to be sent back to requester
		GenericError.Error = "The service is currently unavailable"
		resErr, _ := json.Marshal(GenericError)
		//Write the response back to requester
		w.Write(resErr)
		return
	}

	//If no errors occur, check if the required fields are present
	if Stock.Symbol != "" && Stock.Interval != "" {
		//Query to send to Alpha Vantage
		query := AlphaVantageURL + "function=TIME_SERIES_INTRADAY&symbol=" + Stock.Symbol + "&interval=" + Stock.Interval + "&apikey=" + key.ApiKey
		if Stock.OutputSize != "" {
			query += "&outputsize=" + Stock.OutputSize
		}
		if Stock.DataType != "" {
			query += "&datatype=" + Stock.DataType
		}

		//Call the AplhaVantage API
		resp, err := http.Get(query)
		if err != nil {
			log.Error("Error from AlphaVantage API", err)

			//Set Content-type & Status to client can read the response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			//Prep Error to be sent back to requester
			GenericError.Error = "The service encountered an error calling the AlphaVantage API"
			resErr, _ := json.Marshal(GenericError)
			//Write the response back to requester
			w.Write(resErr)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//Set Content-type & Status to client can read the response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			//Prep Error to be sent back to requester
			GenericError.Error = "The service encountered an error reading the AlphaVantage API Response"
			resErr, _ := json.Marshal(GenericError)
			//Write the response back to requester
			w.Write(resErr)
		}

		//TODO not assigning/outputting the correct Data
		//Assigns the API key to the struct
		err = json.Unmarshal(body, &returnValue)
		if err != nil {
			log.Error("Failed Load API Key to Struct. ", err)

			//Set Content-type & Status to client can read the response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			//Prep Error to be sent back to requester
			GenericError.Error = "The service is currently unavailable"
			resErr, _ := json.Marshal(GenericError)
			//Write the response back to requester
			w.Write(resErr)
			return
		}

		log.Info(string(body))
		log.WithFields(log.Fields{"Info Received": &returnValue}).Info()

	} else {
		log.Error("Symbol or Interval not Set, cannot send query.")

		//Set Content-type & Status to client can read the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		//Prep Error to be sent back to requester
		GenericError.Error = "The service encountered an error returning the data"
		resErr, _ := json.Marshal(GenericError)
		//Write the response back to requester
		w.Write(resErr)
	}

	//Prep Data to be sent back to requester
	Response, err := json.Marshal(returnValue)
	if err != nil {
		log.Error("Failed to Prepare data for send back. ", err)

		//Set Content-type & Status to client can read the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		//Prep Error to be sent back to requester
		GenericError.Error = "The service encountered an error returning the data"
		resErr, _ := json.Marshal(GenericError)
		//Write the response back to requester
		w.Write(resErr)
	}

	//Set Content-type & Status to client can read the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write the response back to requester
	w.Write(Response)

}
