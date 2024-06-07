package server

import (
	"net/http"
)

// Define home handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from this application"))
}
