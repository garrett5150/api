package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"description"`
}

func main() {
	//updates the log file to the current date
	CurrentTime := time.Now().Format("01-02-2006")
	logPath := "./Logs/" + CurrentTime + ".txt"
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error Opening Log File: ", err.Error())
		os.Exit(1)
	}
	//defer f.Close()
	log.SetOutput(f)
	//opens a new router
	router := mux.NewRouter().StrictSlash(true)
	log.Info("Router Successfully Opened")

	//passes the router to routes.go to direct traffic as needed
	routes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
	err = f.Close()
	if err != nil {
		log.Fatal("Failed to close log File")
	}
}

//initiates the Logging routine
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	//Sets the log to TraceLevel severity and above while building the app
	log.SetLevel(log.TraceLevel)
	//opens a log file to track issues, if file doesnt exist, create it.
}
