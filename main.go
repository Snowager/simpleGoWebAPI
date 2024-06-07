package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Snowager/simpleGoWebAPI/server"
)

func main() {
	port := 4000
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.Home)

	//log server start
	log.Printf("Starting server on %v", port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), mux)
	log.Fatal(err)
}
