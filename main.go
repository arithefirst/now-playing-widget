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
	http.HandleFunc("/api/v1/get/STC", getSTC)
	http.HandleFunc("/api/v1/get/TC", getTC)
	http.HandleFunc("/api/v1/get/BG", getBG)

	// Set the functions for the set endpoints
	http.HandleFunc("/api/v1/set/STC", setSTC)
	http.HandleFunc("/api/v1/set/TC", setTC)
	http.HandleFunc("/api/v1/set/BG", setBG)

	// Start the server
	fmt.Printf("Server started on port %v\n", port)
	var portString string = ":" + strconv.Itoa(int(port))
	http.ListenAndServe(portString, nil)
}
