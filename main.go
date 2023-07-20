package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

}
func formHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root.
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is
	// a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 3000\n")
	// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle
	//requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
