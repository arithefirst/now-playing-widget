package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type jsonGet struct {
	Hex string `json:"hex"`
	Err string `json:"err"`
}

// Function to create the json responses for the getX() functions
func sendJsonGet(hex string, err string) jsonGet {
	if hex == "null" {
		json := jsonGet{Err: err, Hex: "null"}
		return json
	} else if err == "null" {
		json := jsonGet{Err: "null", Hex: hex}
		return json
	} else {
		json := jsonGet{Hex: hex, Err: err}
		return json
	}
}

func getBG(c *gin.Context) {
	// Set the content-type header to json and uid querystring var
	c.Header("content-type", "application/json; charset=utf-8")
	uid := c.Query("uid")

	if uid == "" {
		c.IndentedJSON(http.StatusOK, sendJsonGet("null", "uid not present in querystring"))
	} else {
		output, err := get(uid)
		if err != nil {
			c.IndentedJSON(http.StatusOK, sendJsonGet("null", err.Error()))
		} else {
			// Check to see if the value is not set
			if output.Empty {
				// Return default if not set
				c.IndentedJSON(http.StatusOK, sendJsonGet("#181A1B", "No value found: Default returned"))
			} else {
				c.IndentedJSON(http.StatusOK, sendJsonGet(output.BG, "null"))
			}
		}
	}

}

func getTC(c *gin.Context) {
	// Set the content-type header to json and uid querystring var
	c.Header("content-type", "application/json; charset=utf-8")
	uid := c.Query("uid")

	if uid == "" {
		c.IndentedJSON(http.StatusOK, sendJsonGet("null", "uid not present in querystring"))
	} else {
		output, err := get(uid)
		if err != nil {
			c.IndentedJSON(http.StatusOK, sendJsonGet("null", err.Error()))
		} else {
			if output.Empty {
				// Return default if not set
				c.IndentedJSON(http.StatusOK, sendJsonGet("#FFFFFF", "No value found: Default returned"))
			} else {
				c.IndentedJSON(http.StatusOK, sendJsonGet(output.TC, "null"))
			}
		}
	}

}

func getSTC(c *gin.Context) {
	// Set the content-type header to json and uid querystring var
	c.Header("content-type", "application/json; charset=utf-8")
	uid := c.Query("uid")

	if uid == "" {
		c.IndentedJSON(http.StatusOK, sendJsonGet("null", "uid not present in querystring"))
	} else {
		output, err := get(uid)
		if err != nil {
			c.IndentedJSON(http.StatusOK, sendJsonGet("null", err.Error()))
		} else {
			if output.Empty {
				// Return default if not set
				c.IndentedJSON(http.StatusOK, sendJsonGet("#D3D3D3", "No value found: Default returned"))
			} else {
				c.IndentedJSON(http.StatusOK, sendJsonGet(output.STC, "null"))
			}
		}
	}

}

func getRight(c *gin.Context) {
	// Set the content-type header to json and uid querystring var
	c.Header("content-type", "application/json; charset=utf-8")
	uid := c.Query("uid")

	if uid == "" {
		c.IndentedJSON(http.StatusOK, sendJsonGet("null", "uid not present in querystring"))
	} else {
		output, err := get(uid)
		if err != nil {
			c.IndentedJSON(http.StatusOK, sendJsonGet("null", err.Error()))
		} else {
			if output.Empty {
				// Return default if not set
				c.IndentedJSON(http.StatusOK, sendJsonGet("false", "No value found: Default returned"))
			} else {
				c.IndentedJSON(http.StatusOK, sendJsonGet(fmt.Sprint(output.RIGHT), "null"))
			}
		}
	}

}
