package main

import (
	"log"
	"net/http"
	"strconv"
)

// Define home handler function
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from this application"))
}

func main() {
	port := 4000
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	//log server start
	log.Printf("Starting server on %v", port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), mux)
	log.Fatal(err)
}
