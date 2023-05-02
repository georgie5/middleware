package main

//Middleware is a way to add functionality to your HTTP server
//such as authentication or logging, and can be used to perform tasks before and after a request is handled by the server.

//The first party middleware is using the goji/httpauth package, which provides HTTP Basic Authentication functionality.
//This package is used by calling a helper function to set up the middleware.
//The middleware function is then used in the same way as any custom-built middleware.

//how to use the httpauth middleware package for HTTP Basic Authentication

// This function takes a destination io.Writer. which is used to write logs to a firl or another output place
// it returns a function that takes an http.Handler as an argument and returns and http.Hander
