package main

import (
	"fmt"
	"net/http"
)

func main() {
	// mux = multiplexer aka handler
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("api request recieved")
		fmt.Println(r.Method)
		w.Write([]byte("Hello World"))
	})

	mux.HandleFunc("/go/", func(w http.ResponseWriter, r *http.Request) {
		// these two lines are printing on server side
		fmt.Println("api request recieved")
		fmt.Println(r.Method)

		// this is the responce to the api client
		w.Write([]byte("Start your engine"))
	})

	// below line is setting up where our server listening for requests
	http.ListenAndServe("localhost:3000", mux)
}
