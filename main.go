package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Just checking if the path is correct
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// Allowing only get method
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello Guys!!")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Just checking if the path is correct
	if r.URL.Path != "/formr" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// Allowing only post method
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Println("Post successful")
	name := r.FormValue("name")
	fmt.Fprintf(w, "Name = %s", name)

}

func main() {
	// FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root.
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is
	// a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/form.html")
	})
	http.HandleFunc("/formr", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 3000\n")
	// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle
	//requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
