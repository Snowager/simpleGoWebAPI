package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := 4000
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	//log server start
	log.Printf("Starting server on %v", port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), mux)
	log.Fatal(err)
}
