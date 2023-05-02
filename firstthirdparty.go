package main

/*
//The first party middleware is using the goji/httpauth package, which provides HTTP Basic Authentication functionality.
//This package is used by calling a helper function to set up the middleware.
//The middleware function is then used in the same way as any custom-built middleware.

//how to use the httpauth middleware package for HTTP Basic Authentication

import (
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

func main() {
	//this line creates an authentication handler that requires a username and password
	authHandler := httpauth.SimpleBasicAuth("alice", "pa$$word1")

	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", authHandler(finalHandler)) //This line adds the auhentication handler as a middleware

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
*/
