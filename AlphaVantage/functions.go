package AlphaVantage

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func returnError(w http.ResponseWriter, message string) {
	var GenericError Error

	//Set Content-type & Status to client can read the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	//Prep Error to be sent back to requester
	GenericError.Error = message
	resErr, _ := json.Marshal(GenericError)
	//Write the response back to requester
	w.Write(resErr)
}
