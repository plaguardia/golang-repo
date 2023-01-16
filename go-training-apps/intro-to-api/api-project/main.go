package main

import (
	"net/http"
)

func main() {

	// mux = multiplexer aka handler
	mux := http.NewServeMux()

	// NOTE
	//	I am not able to use local packages like "structs" in this case in a folder outside of ~/go due to the gopath setting
	// I am jumping over there to test some things but might need to move me repo over there for local testing to move forward with my learning

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.responce
		}

	})

	// below line is setting up where our server listening for requests
	http.ListenAndServe("localhost:3000", mux)
}
