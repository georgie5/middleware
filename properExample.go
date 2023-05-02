package main

/*
import (
	"log"
	"mime"
	"net/http"
)

//We want to create some middleware that checks for the existence of a Content-Type header and verifies that the header's value is "application/json."
//If either of these checks fails, we want our middleware to write an error message and stop the request from reaching our application handlers

//The enforceJSONHandler function is the middleware that we're building. It takes an http.Handler as input and returns an http.Handler.
//This function is responsible for intercepting the incoming requests, checking the Content-Type header, and ensuring that it has the correct value.

func enforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		//Inside the enforceJSONHandler function, we first extract the Content-Type header from the incoming request using r.Header.Get("Content-Type").
		//If the Content-Type header is not empty, we proceed to the next step.

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)

			// If the function returns an error, it means that the Content-Type header is malformed, and we return an HTTP error response with the status code of 400 (Bad Request).
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}

			//if the mime type returned by the function is not "application/json," we return an HTTP error response with the status code of 415
			//(Unsupported Media Type).

			//If both checks pass, we call next.ServeHTTP(w, r) to pass the request to the next middleware or handler in the chain.

			if mt != "application/json" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

//The final handler just writes "OK" as a respons
//In the main function, we create a new ServeMux and register our final handler using the mux.Handle("/", enforceJSONHandler(finalHandler)) statement.

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", enforceJSONHandler(finalHandler))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

/*
To test our middleware, we use curl to make requests to our service. In the first request,
 we don't include the Content-Type header, and our middleware returns a bad Request response.
 In the second request, we include the Content-Type header with the value application/xml
 and our middleware returns a 415 Unsupported Media Type response.
 In the third request, we include the Content-Type header with the value "application/json; charset=UTF-8," and our final handler responds with "OK."
*/
