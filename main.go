package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// Set the port for the server to run on
	var port uint16 = 5500

	// Set the functions for the get endpoints
	http.HandleFunc("/api/v1/get/stc", getSTC)
	http.HandleFunc("/api/v1/get/tc", getTC)
	http.HandleFunc("/api/v1/get/bg", getBG)

	// Set the functions for the set endpoints
	http.HandleFunc("/api/v1/set/stc", setSTC)
	http.HandleFunc("/api/v1/set/tc", setTC)
	http.HandleFunc("/api/v1/set/bg", setBG)

	// Serve the files in static/
	http.Handle("/", http.FileServer(http.Dir("static/")))
	http.Handle("/config/", http.StripPrefix("/config/", http.FileServer(http.Dir("static/config/"))))

	// Start the server
	fmt.Printf("Server started on port %v\n", port)
	var portString string = ":" + strconv.Itoa(int(port))
	http.ListenAndServe(portString, nil)
}
