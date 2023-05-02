package main

/*
//The second party is the gorilla/handlers package, which provides a LoggingHandler middleware that records request logs using the Apache Common Log Format.
//This middleware takes not only the next handler but also the io.Writer that the log will be written to.
//In this example, the log is written to a server.log file in the current directory.

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	//This line creates a new file called "server.log" (or opens it if it already exists) with write-only
	// creates it if it doesn't exist, and appends to it if it already exists.
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	//This line adds the LoggingHandler middleware to the default multiplexer mux, with finalHandler as the final handler to be called in the chain.
	finalHandler := http.HandlerFunc(final)
	//The LoggingHandler middleware records the request logs and writes them to the logFile using the Apache Common Log Format.
	//The finalHandler writes the "OK" response
	mux.Handle("/", handlers.LoggingHandler(logFile, finalHandler))

	log.Print("Listening on :3000...")
	err = http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

/*
Fow now the code might look understandable but when using multiple middleware the code can quickly become difficult to read and will be confusing
*/
