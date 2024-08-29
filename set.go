package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Function to create the json responses for the setConf() func
func sendJsonSet(upsert int64, mod int64, err string) string {
	if err == "null" {
		json := fmt.Sprintf("{\"inserted\":%d,\"modified\":%d,\"err\":null}", upsert, mod)
		return json
	} else {
		json := fmt.Sprintf("{\"inserted\":%d,\"modified\":%d,\"err\":\"%v\"}", upsert, mod, err)
		return json
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

func setConf(w http.ResponseWriter, r *http.Request) {
	// Set the content-type header to json and utf-8
	w.Header().Set("content-type", "application/json; charset=utf-8")

	// Check to see if bg, tc, stc, right, and uid are present in the querystring
	if r.URL.Query().Get("bg") == "" ||
		r.URL.Query().Get("tc") == "" ||
		r.URL.Query().Get("stc") == "" ||
		r.URL.Query().Get("right") == "" ||
		r.URL.Query().Get("uid") == "" {
		fmt.Fprint(w, "{\"err\":\"All inputs are required to use this endpoint\"}")
	} else {
		// Grab the auth header from the request
		auth := r.Header.Get("Authorization")

		// If the uid response from the authCheck() function is
		// not the same as the user given UID, return an error
		if r.URL.Query().Get("uid") != authCheck(auth) {
			// Handling for each possible type of error
			if strings.Contains(authCheck(auth), "No token provided") {
				fmt.Fprint(w, "{\"err\":\"No Token Provided\"}")
			} else if strings.Contains(authCheck(auth), "The access token expired") {
				fmt.Fprint(w, "{\"err\":\"The access token expired\"}")
			} else if strings.Contains(authCheck(auth), "Only valid bearer authentication supported") {
				fmt.Fprint(w, "{\"err\":\"Only valid bearer authentication supported\"}")
			} else {
				fmt.Fprint(w, "{\"err\":\"You cannot modify another user's configuration\"}")
			}
		} else {
			// Default the alignRight var to false, set to true if
			// the value of "right" in the querystring is "true"
			var alignRight bool = false
			if r.URL.Query().Get("right") == "true" {
				alignRight = true
			}

			// Populate the user struct with the paramaters from
			// the querystring
			updateParams := user{
				BG:    r.URL.Query().Get("bg"),
				TC:    r.URL.Query().Get("tc"),
				STC:   r.URL.Query().Get("stc"),
				RIGHT: alignRight,
				UID:   r.URL.Query().Get("uid"),
			}

			// Call the function to udpate the database with the
			// paramaters from the querystring

			result, err := set(updateParams)
			if err != nil {
				fmt.Fprint(w, sendJsonSet(0, 0, err.Error()))
			} else {
				// Return the number of docs modified or inserted
				fmt.Fprint(w, sendJsonSet(result.UpsertedCount, result.ModifiedCount, "null"))
			}
		}
	}

}
