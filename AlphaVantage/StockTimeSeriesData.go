package AlphaVantage

//https://www.alphavantage.co/documentation/

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type StockReq struct {
	Symbol     string `json:"symbol"`
	Interval   string `json:"interval,omitempty"`
	OutputSize string `json:"outputsize,omitempty"`
	DataType   string `json:"datatype,omitempty"`
}

type ApiKey struct {
	ApiKey string `json:"apikey"`
}

//TODO Incorrect incoming request such as http://localhost:8080/AlphaVantage/STS/MonthlyAdjustedasdsadasdsadasda breaks horribly rather than returning a 400 error
//TODO might be a int related issue

var AlphaVantageURL = "https://www.alphavantage.co/query?"

func StockTimeSeries(w http.ResponseWriter, r *http.Request) {
	//Grab the name of the function to be called
	vars := mux.Vars(r)
	function := vars["NAME"]

	log.Trace(function + " Called")
	var returnValue interface{}
	var Stock StockReq
	var key ApiKey

	//reads the JSON from the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read Body. ", err)
		returnError(w, "The server could not understand the request due to invalid syntax", 400)
	}
	//Assigns JSON Request to the Stock Structure
	err = json.Unmarshal(body, &Stock)
	if err != nil {
		log.Error("Failed Load Data to Stock Struct. ", err)
		returnError(w, "The server could not understand the request due to invalid syntax", 400)
	}
	//Log the info we received
	log.WithFields(log.Fields{"Info Received": &Stock}).Info()

	//grabs the Alpha Vantage API Key
	file, err := ioutil.ReadFile("./AlphaVantage/AlphaVantageApiKey.json")
	if err != nil {
		log.Error("Failed to read ApiKey File. ", err)
		returnError(w, "The service is currently unavailable", 502)
	}
	//Assigns the API key to the struct
	err = json.Unmarshal(file, &key)
	if err != nil {
		log.Error("Failed Load API Key to Struct. ", err)
		returnError(w, "The service is currently unavailable", 503)
	}

	//If no errors occur, check if the required fields are present
	if Stock.Symbol != "" {
		//Calls the relevant function to build the proper query
		query := ""
		switch function {
		case "Intraday":
			if Stock.Interval != "" {
				query = TimeSeriesIntraday(Stock, key)
			} else {
				log.Error("Interval not Set, cannot send query.")
				returnError(w, "The service encountered an error returning the data", 400)
			}
		case "Daily":
			query = TimeSeriesDaily(Stock, key)
		case "DailyAdjusted":
			query = TimeSeriesDailyAdjusted(Stock, key)
		case "Weekly":
			query = TimeSeriesWeekly(Stock, key)
		case "WeeklyAdjusted":
			query = TimeSeriesWeeklyAdjusted(Stock, key)
		case "Monthly":
			query = TimeSeriesMonthly(Stock, key)
		case "MonthlyAdjusted":
			query = TimeSeriesMonthlyAdjusted(Stock, key)
		case "QuoteEndpoint":
			query = QuoteEndpoint(Stock, key)
		case "SearchEndpoint":
			query = SearchEndpoint(Stock, key)
		default:
			{
				log.Error("Unknown switch function call")
				returnError(w, function+" is an unknown value for STS calls", 400)
			}
		}

		//Call the AplhaVantage API
		resp, err := http.Get(query)
		if err != nil {
			log.Error("Error from AlphaVantage API", err)
			returnError(w, "The service encountered an error calling the AlphaVantage API", 502)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error("Failed to read AlphaVantage API Response")
			returnError(w, "The service encountered an error reading the AlphaVantage API Response", 500)
		}

		//Assigns the API key to the struct
		err = json.Unmarshal(body, &returnValue)
		if err != nil {
			log.Error("Failed Load API Key to Struct. ", err)
			returnError(w, "The service is currently unavailable", 500)
		}

	} else {
		log.Error("Symbol not Set, cannot send query.")
		returnError(w, "The service encountered an error returning the data", 400)
	}

	//Prep Data to be sent back to requester
	Response, err := json.Marshal(returnValue)
	if err != nil {
		log.Error("Failed to Prepare data for send back. ", err)
		returnError(w, "The service encountered an error returning the data", 500)
	}

	//Set Content-type & Status to client can read the response
	log.Info("Query Successful, returning: ", returnValue)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write the response back to requester
	w.Write(Response)
	log.Trace("Closing " + function + " function")
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
func TimeSeriesIntraday(Stock StockReq, key ApiKey) string {
	//check if the required fields are present
	//build the query to be sent to Alpha Vantage
	query := AlphaVantageURL + "function=TIME_SERIES_INTRADAY&symbol=" + Stock.Symbol + "&interval=" + Stock.Interval
	if Stock.OutputSize != "" {
		query += "&outputsize=" + Stock.OutputSize
	}
	if Stock.DataType != "" {
		query += "&datatype=" + Stock.DataType
	}
	//log the query for debugging before adding the api key
	log.Debug("Query to be sent: " + query)
	query += "&apikey=" + key.ApiKey
	return query
}

//This API returns daily time series (date, daily open, daily high, daily low, daily close, daily volume) of the global equity specified, covering 20+ years of historical data.
//Required Parameters
//function=TIME_SERIES_DAILY
//symbol=IBM
//apikey=key
//Optional Parameters
//outputsize=compact, full
//datatype=json, csv
func TimeSeriesDaily(Stock StockReq, key ApiKey) string {
	query := ""
	//Query to send to Alpha Vantage
	query = AlphaVantageURL + "function=TIME_SERIES_DAILY&symbol=" + Stock.Symbol
	if Stock.OutputSize != "" {
		query += "&outputsize=" + Stock.OutputSize
	}
	if Stock.DataType != "" {
		query += "&datatype=" + Stock.DataType
	}
	//log the query for debugging before adding the api key
	log.Debug("Query to be sent: " + query)
	query += "&apikey=" + key.ApiKey

	return query
}

//This API returns daily time series (date, daily open, daily high, daily low, daily close, daily volume, daily adjusted close, and split/dividend events) of the global equity specified, covering 20+ years of historical data.
//Required Parameters
//function=TIME_SERIES_DAILY
//symbol=IBM
//apikey=key
//Optional Parameters
//outputsize=compact, full
//datatype=json, csv
func TimeSeriesDailyAdjusted(Stock StockReq, key ApiKey) string {
	//Query to send to Alpha Vantage
	query := AlphaVantageURL + "function=TIME_SERIES_DAILY_ADJUSTED&symbol=" + Stock.Symbol
	if Stock.OutputSize != "" {
		query += "&outputsize=" + Stock.OutputSize
	}
	if Stock.DataType != "" {
		query += "&datatype=" + Stock.DataType
	}
	//log the query for debugging before adding the api key
	log.Debug("Query to be sent: " + query)
	query += "&apikey=" + key.ApiKey
	return query
}

//This API returns weekly time series (last trading day of each week, weekly open, weekly high, weekly low, weekly close, weekly volume) of the global equity specified, covering 20+ years of historical data.
//Required Parameters
//function=TIME_SERIES_WEEKLY
//symbol=IBM
//apikey=key
//Optional Parameters
//datatype=json, csv
func TimeSeriesWeekly(Stock StockReq, key ApiKey) string {

	query := AlphaVantageURL + "function=TIME_SERIES_WEEKLY&symbol=" + Stock.Symbol
	if Stock.DataType != "" {
		query += "&datatype=" + Stock.DataType
	}
	//log the query for debugging before adding the api key
	log.Debug("Query to be sent: " + query)
	query += "&apikey=" + key.ApiKey
	return query
}

//This API returns weekly adjusted time series (last trading day of each week, weekly open, weekly high, weekly low, weekly close, weekly adjusted close, weekly volume, weekly dividend) of the global equity specified, covering 20+ years of historical data.
//Required Parameters
//function=TIME_SERIES_WEEKLY_ADJUSTED
//symbol=IBM
//apikey=key
//Optional Parameters
//datatype=json, csv
func TimeSeriesWeeklyAdjusted(Stock StockReq, key ApiKey) string {
	//Query to send to Alpha Vantage
	query := AlphaVantageURL + "function=TIME_SERIES_WEEKLY_ADJUSTED&symbol=" + Stock.Symbol
	if Stock.DataType != "" {
		query += "&datatype=" + Stock.DataType
	}
	//log the query for debugging before adding the api key
	log.Debug("Query to be sent: " + query)
	query += "&apikey=" + key.ApiKey
	return query
}

//This API returns monthly time series (last trading day of each month, monthly open, monthly high, monthly low, monthly close, monthly volume) of the global equity specified, covering 20+ years of historical data.
//Required Parameters
//function=TIME_SERIES_MONTHLY
//symbol=IBM
//apikey=key
//Optional Parameters
//datatype=json, csv
func TimeSeriesMonthly(Stock StockReq, key ApiKey) string {
	//Query to send to Alpha Vantage
	query := AlphaVantageURL + "function=TIME_SERIES_MONTHLY&symbol=" + Stock.Symbol
	if Stock.DataType != "" {
		query += "&datatype=" + Stock.DataType
	}
	//log the query for debugging before adding the api key
	log.Debug("Query to be sent: " + query)
	query += "&apikey=" + key.ApiKey
	return query
}

//This API returns monthly adjusted time series (last trading day of each month, monthly open, monthly high, monthly low, monthly close, monthly adjusted close, monthly volume, monthly dividend) of the equity specified, covering 20+ years of historical data.
//Required Parameters
//function=TIME_SERIES_MONTHLY_ADJUSTED
//symbol=IBM
//apikey=key
//Optional Parameters
//datatype=json, csv
func TimeSeriesMonthlyAdjusted(Stock StockReq, key ApiKey) string {
	//Query to send to Alpha Vantage
	query := AlphaVantageURL + "function=TIME_SERIES_MONTHLY_ADJUSTED&symbol=" + Stock.Symbol
	if Stock.DataType != "" {
		query += "&datatype=" + Stock.DataType
	}
	//log the query for debugging before adding the api key
	log.Debug("Query to be sent: " + query)
	query += "&apikey=" + key.ApiKey
	return query
}

//A lightweight alternative to the time series APIs, this service returns the latest price and volume information for a security of your choice.
//Required Parameters
//function=GLOBAL_QUOTE
//symbol=IBM
//apikey=key
//Optional Parameters
//datatype=json, csv
func QuoteEndpoint(Stock StockReq, key ApiKey) string {

	//Query to send to Alpha Vantage
	query := AlphaVantageURL + "function=GLOBAL_QUOTE&symbol=" + Stock.Symbol
	if Stock.DataType != "" {
		query += "&datatype=" + Stock.DataType
	}
	//log the query for debugging before adding the api key
	log.Debug("Query to be sent: " + query)
	query += "&apikey=" + key.ApiKey
	return query
}

//The Search Endpoint returns the best-matching symbols and market information based on keywords of your choice. The search results also contain match scores that provide you with the full flexibility to develop your own search and filtering logic.
//Required Parameters
//function=GLOBAL_QUOTE
//keywords=IBM
//apikey=key
//Optional Parameters
//datatype=json, csv
func SearchEndpoint(Stock StockReq, key ApiKey) string {

	//Query to send to Alpha Vantage
	query := AlphaVantageURL + "function=SYMBOL_SEARCH&keywords=" + Stock.Symbol
	if Stock.DataType != "" {
		query += "&datatype=" + Stock.DataType
	}
	//log the query for debugging before adding the api key
	log.Debug("Query to be sent: " + query)
	query += "&apikey=" + key.ApiKey
	return query
}
