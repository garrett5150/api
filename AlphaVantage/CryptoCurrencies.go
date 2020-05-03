package AlphaVantage

import "net/http"

//This API returns the realtime exchange rate for any pair of digital currency (e.g., Bitcoin) or physical currency (e.g., USD).
//Required Parameters
//function=CURRENCY_EXCHANGE_RATE
//from_currency=USD
//to_currency=CAD
//apikey=key
func CryptoExchangeRates(w http.ResponseWriter, r *http.Request) {

}

//Fundamental Crypto Asset Score (FCAS) is a comparative metric used to assess the fundamental health of crypto projects. The score is derived from the interactivity between primary project life-cycle factors: User Activity/Utility, Developer Behavior, and Market Maturity. Each crypto asset is given a composite numerical score, 0-1000, and an associated rating as follows:
//Superb 900-1000
//Attractive 750-899
//Basic 650-749
//Caution 500-649
//Fragile Below 500

//Required Parameters
//function=CRYPTO_RATING
//symbol=BTC
//apikey=key
func CryptoHealthIndex(w http.ResponseWriter, r *http.Request) {

}

//This API returns the daily historical time series for a digital currency (e.g., BTC) traded on a specific market (e.g., CNY/Chinese Yuan), refreshed daily at midnight (UTC). Prices and volumes are quoted in both the market-specific currency and USD.
//Required Parameters
//function=DIGITAL_CURRENCY_DAILY
//symbol=BTC
//market=CNY
//apikey=key
func CryptoDaily(w http.ResponseWriter, r *http.Request) {

}

//This API returns the weekly historical time series for a digital currency (e.g., BTC) traded on a specific market (e.g., CNY/Chinese Yuan), refreshed daily at midnight (UTC). Prices and volumes are quoted in both the market-specific currency and USD
//Required Parameters
//function=DIGITAL_CURRENCY_WEEKLY
//symbol=BTC
//market=CNY
//apikey=key
func CryptoWeekly(w http.ResponseWriter, r *http.Request) {

}

//This API returns the monthly historical time series for a digital currency (e.g., BTC) traded on a specific market (e.g., CNY/Chinese Yuan), refreshed daily at midnight (UTC). Prices and volumes are quoted in both the market-specific currency and USD.
//Required Parameters
//function=DIGITAL_CURRENCY_MONTHLY
//symbol=BTC
//market=CNY
//apikey=key
func CryptoMonthly(w http.ResponseWriter, r *http.Request) {

}
