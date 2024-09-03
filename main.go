package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the port for the server to run on
	var port uint16 = 5500

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Set the functions for the get endpoints
	router.GET("/api/v1/get/stc", getSTC)
	router.GET("/api/v1/get/tc", getTC)
	router.GET("/api/v1/get/bg", getBG)
	router.GET("/api/v1/get/right", getRight)

	// Set the functions for the set endpoint
	router.GET("/api/v1/set", setConf)

	// Serve the files in static/
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/main.js", "./static/main.js")
	router.StaticFile("/styles.css", "./static/styles.css")
	router.StaticFile("/jquery-3.7.1.min.js", "./static/jquery-3.7.1.min.js")

	router.Static("/config", "./static/config")

	// Start the server
	fmt.Printf("Server started on port %v\n", port)
	err := router.Run(fmt.Sprintf("127.0.0.1:%v", port))
	if err != nil {
		panic(err)
	}
}
