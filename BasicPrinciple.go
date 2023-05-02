package main

/*
import "net/http"

//What's happening here is that we're defining a function that takes a message string as a parameter.
//This function then returns a handler that uses the http.HandlerFunc.
//It writes the message string to the response into an http.Handler.
//This http.Handler is then returned by the messageHandler function and can be used to handle incoming HTTP requests.

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})
}

// This function creates a middleware that will be inserted into a chain of handlers.

func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		next.ServeHTTP(w, r)
	})
}

/*
What's happening here is that we're defining a function that takes an http.Handler called next as a parameter.
This next handler represents the next handler in the chain. Inside the function
we're defining a closure that will act as our middleware logic.
This closure takes in the http.ResponseWriter and *http.Request and then calls the next.ServeHTTP(w, r) function
and what thiss does is that it passes control to the next handler in the chain.
So this middleware function allows us to perform certain actions on the request
and response before passing it on to the next handler in the chain.
*/
