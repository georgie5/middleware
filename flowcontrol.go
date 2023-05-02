package main

/*
import (
	"log"
	"net/http"
)

//handlers: middlewareOne, middlewareTwo, and final.
//These handlers form a handler chain that will be used to handle incoming HTTP requests.

//When a request comes in, it is first passed to middlewareOne, which logs a message to the console and then passes control to the next handler in the chain
//by calling next.ServeHTTP(w, r).

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Print("Executing middlewareOne again")
	})
}

//This next handler is middlewareTwo, which also logs a message to the console and checks if the requested URL path is "/foo".
//If it is, middlewareTwo stops the flow of control by returning immediately. Otherwise, it passes control to the final handler,
//which logs a message to the console and writes the "OK" string to the response body.

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareTwo")
		if r.URL.Path == "/foo" {
			return
		}

		next.ServeHTTP(w, r)
		log.Print("Executing middlewareTwo again")
	})
}

//After the final handler completes the control is passed back up the chain in reverse order.
//This means that middlewareTwo is executed again and it logs another message to the console before passing control back to middlewareOne.
//Finally middlewareOne logs a final message to the console before returning control to the web server.

func final(w http.ResponseWriter, r *http.Request) {
	log.Print("Executing finalHandler")
	w.Write([]byte("OK"))
}

//If we were to visit the URL http://localhost:3000/foo instead, middlewareTwo would stop the flow of control by returning immediately,
//and neither the final handler nor any of the subsequent middleware would be executed.

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", middlewareOne(middlewareTwo(finalHandler)))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
*/
