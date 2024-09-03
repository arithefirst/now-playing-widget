package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type jsonSet struct {
	upsert int64
	mod    int64
	err    string
}

// Function to create the json responses for the setConf() func
func sendJsonSet(upsert int64, mod int64, err string) jsonSet {
	if err == "null" {
		jsonReturn := jsonSet{upsert: upsert, mod: mod, err: "null"}
		return jsonReturn
	} else {
		jsonReturn := jsonSet{upsert: upsert, mod: mod, err: err}
		return jsonReturn
	}
}

func authCheck(auth string) string {
	client := &http.Client{}

	// Prepare http request
	req, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return err.Error()
	}

	// Add the auth header to the http request
	req.Header.Set("Authorization", auth)

	// Send request
	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}

	// Read request body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	// Convert the json body into a map
	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return err.Error()
	}

	// If the "error" field is present, return the whole body,
	// otherwise return the ID field
	_, errorExists := result["error"]
	if errorExists {
		return string(body)
	} else {
		id, ok := result["id"].(string)
		if !ok {
			return "Failed to parse json value"
		} else {
			return id
		}
	}

}

func setConf(c *gin.Context) {
	// Set the content-type header to json and utf-8
	c.Header("content-type", "application/json; charset=utf-8")

	// Check to see if bg, tc, stc, right, and uid are present in the querystring
	if c.Query("bg") == "" ||
		c.Query("tc") == "" ||
		c.Query("stc") == "" ||
		c.Query("right") == "" ||
		c.Query("uid") == "" {
		c.IndentedJSON(http.StatusOK, "{\"err\":\"All inputs are required to use this endpoint\"}")
	} else {
		// Grab the auth header from the request
		auth := c.GetHeader("Authorization")

		// If the uid response from the authCheck() function is
		// not the same as the user given UID, return an error
		if c.Query("uid") != authCheck(auth) {
			// Handling for each possible type of error
			if strings.Contains(authCheck(auth), "No token provided") {
				c.IndentedJSON(http.StatusOK, "{\"err\":\"No Token Provided\"}")
			} else if strings.Contains(authCheck(auth), "The access token expired") {
				c.IndentedJSON(http.StatusOK, "{\"err\":\"The access token expired\"}")
			} else if strings.Contains(authCheck(auth), "Only valid bearer authentication supported") {
				c.IndentedJSON(http.StatusOK, "{\"err\":\"Only valid bearer authentication supported\"}")
			} else {
				c.IndentedJSON(http.StatusOK, "{\"err\":\"You cannot modify another user's configuration\"}")
			}
		} else {
			// Default the alignRight var to false, set to true if
			// the value of "right" in the querystring is "true"
			var alignRight bool = false
			if c.Query("right") == "true" {
				alignRight = true
			}

			// Populate the user struct with the parameters from
			// the querystring
			updateParams := user{
				BG:    c.Query("bg"),
				TC:    c.Query("tc"),
				STC:   c.Query("stc"),
				RIGHT: alignRight,
				UID:   c.Query("uid"),
			}

			// Call the function to update the database with the
			// parameters from the querystring

			result, err := set(updateParams)
			if err != nil {
				c.IndentedJSON(http.StatusOK, sendJsonSet(0, 0, err.Error()))
			} else {
				// Return the number of docs modified or inserted
				c.IndentedJSON(http.StatusOK, sendJsonSet(result.UpsertedCount, result.ModifiedCount, "null"))
			}
		}
	}

}
