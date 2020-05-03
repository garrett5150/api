package AlphaVantage

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

//TODO Incorrect incoming request such as http://localhost:8080/AlphaVantage/STS/MonthlyAdjustedasdsadasdsadasda breaks horribly rather than returning a 400 error
//TODO might be a int related issue

//Generic error function that returns an error code, and attached string
func returnError(w http.ResponseWriter, message string, status int) {
	var GenericError Error

	//Set Content-type & Status to client can read the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	//Prep Error to be sent back to requester
	GenericError.Error = message
	resErr, _ := json.Marshal(GenericError)
	//Write the response back to requester
	w.Write(resErr)
}
