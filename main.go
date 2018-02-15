package main

import (
	"encoding/json"
	"log"
	"net/http"
	"utilities"
)

var store map[DuckTypes.ID]DuckTypes.Message

func main() {

	http.HandleFunc("/post", PostAdviceEndpoint)
	http.HandleFunc("/respond", RespondToPostEndpoint)

	// Serve the above REST endpoints on localhost on port 8080
	err := http.ListenAndServe(":8080", nil)

	// If something's broken, fail!
	if err != nil {
		log.Fatal(err)
	}
}

// Some simple API functions

func PostAdviceEndpoint(res http.ResponseWriter, req *http.Response) {
	// Generate a new post from the JSON in the response

	// Commit that post to the store

	// Write a response back to the client
	res.Write([]byte(`{"success": -1, "error": "Endpoint not implemented!"}`))
}

func RespondToPostEndpoint(res http.ResponseWriter, req *http.Response) {
	// Generate a new rubberduck response value from the json in the http response

	// Log that response against the relevant message in the store

	// Write a response back to the client
	res.Write([]byte(`{"success": -1, "error": "Endpoint not implemented!"}`))
}

func ViewPostsEndpoint(res http.ResponseWriter, req *http.Response) {
	// Convert the store to JSON appropriate to send the client for the various posts

	// Send that JSON
	res.Write([]byte(`{"success": -1, "error": "Endpoint not implemented!"}`))
}

func ViewResponsesEndpoint(res http.ResponseWriter, req *http.Response) {
	// Format the responses for the logged in user, so they can be sent to the client

	// Render the responses inside a template
	res.Write([]byte(`{"success": -1, "error": "Endpoint not implemented!"}`))
}

func LoginEndpoint(res http.ResponseWriter, req *http.Request) {
	// Check login details

	// Generate and return an auth token
	res.Write([]byte(`{"success": -1, "error": "Endpoint not implemented!"}`))
}

func LogoutHandler(res http.ResponseWriter, req *http.Request) {
	// Remove relevant auth token

	// Return logout success
	res.Write([]byte(`{"success": -1, "error": "Endpoint not implemented!"}`))
}
