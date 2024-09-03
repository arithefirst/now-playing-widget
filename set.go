package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type jsonSet struct {
	Upsert int64 `json:"upsert"`
	Mod    int64 `json:"mod"`
	Err    error `json:"err"`
	Values user  `json:"values"`
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
	c.Header("content-type", "application/json; charset=utf-8")

	type requestBody struct {
		UID   string `json:"uid"`
		BG    string `json:"bg"`
		TC    string `json:"tc"`
		STC   string `json:"stc"`
		RIGHT bool   `json:"right"`
	}

	type returnErr struct {
		Err string `json:"err"`
	}

	var postRequest requestBody
	err := c.BindJSON(&postRequest)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// Check to see if any of the values are empty
	if postRequest.UID == "" || postRequest.BG == "" || postRequest.STC == "" || postRequest.TC == "" {
		err := returnErr{Err: "All fields must be present"}
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	// Get the current user's UID from the spotify
	// API and ensure it matches the specified UID
	auth := c.GetHeader("Authorization")
	if postRequest.UID != authCheck(auth) {
		// Handling for each possible type of error
		if strings.Contains(authCheck(auth), "No token provided") {
			err := returnErr{Err: "No token provided"}
			c.IndentedJSON(http.StatusUnauthorized, err)
			return
		} else if strings.Contains(authCheck(auth), "The access token expired") {
			err := returnErr{Err: "The access token expired"}
			c.IndentedJSON(http.StatusUnauthorized, err)
			return
		} else if strings.Contains(authCheck(auth), "Only valid bearer authentication supported") {
			err := returnErr{Err: "Only valid bearer authentication supported"}
			c.IndentedJSON(http.StatusUnauthorized, err)
			return
		} else {
			err := returnErr{Err: "You cannot modify another user's configuration"}
			c.IndentedJSON(http.StatusUnauthorized, err)
			return
		}
	}

	updateParams := user{
		UID:   postRequest.UID,
		BG:    postRequest.BG,
		TC:    postRequest.TC,
		STC:   postRequest.STC,
		RIGHT: postRequest.RIGHT,
	}

	result, err := set(updateParams)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, jsonSet{Upsert: 0, Mod: 0, Err: err, Values: updateParams})
	} else {
		// Return the number of docs modified or inserted
		c.IndentedJSON(http.StatusOK, jsonSet{Upsert: result.UpsertedCount, Mod: result.ModifiedCount, Err: nil, Values: updateParams})
	}

}
