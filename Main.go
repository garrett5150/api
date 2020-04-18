package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"description"`
}

type error interface {
	Error() string
}

/*var standardError = log.WithFields(log.Fields{
"description": "",
"Error":       "",
})*/

func main() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error Opening Log File: ", err.Error())
	}
	defer f.Close()
	log.SetOutput(f)
	//opens a new router
	router := mux.NewRouter().StrictSlash(true)
	log.Info("Router Successfully Opened")

	//passes the router to routes.go to direct traffic as needed
	routes(router)
	log.Fatal(http.ListenAndServe(":8080", router))

}

//initiates the Logging routine
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	//Sets the log to TraceLevel severity and above while building the app
	log.SetLevel(log.TraceLevel)
	//opens a log file to track issues, if file doesnt exist, create it.
}

//test Route functions
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!")
}
func test1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test1!")
}
func test2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test2!")
}
