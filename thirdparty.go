package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// This function takes a destination io.Writer. which is used to write logs to a firl or another output place
// it returns a function that takes an http.Handler as an argument and returns and http.Hander
func newLoggingHandler(dst io.Writer) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		// The returned http handler is the LogginHandler middleware that the gorilla handler package provides
		return handlers.LoggingHandler(dst, h)
	}
}

// By using this constructor function, it can now be easy to add to logging middlware to the request handling chain
func main() {
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}
	// Create a new logging handler using the log file
	loggingHandler := newLoggingHandler(logFile)

	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", loggingHandler(finalHandler))

	log.Print("Listening on :3000...")
	err = http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

// The final request handler that returns "OK"
func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
