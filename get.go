package main

import (
	"fmt"
	"net/http"
)

// Function to create the json responses for the getX() funcs
func sendJsonGet(hex string, err string) string {
	if hex == "null" {
		json := fmt.Sprintf("{\"hex\":%v,\"err\":\"%v\"}", "null", err)
		return json
	} else if err == "null" {
		json := fmt.Sprintf("{\"hex\":\"%v\",\"err\":%v}", hex, "null")
		return json
	} else {
		json := fmt.Sprintf("{\"hex\":\"%v\",\"err\":\"%v\"}", hex, err)
		return json
	}
}

func getBG(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	uid := r.URL.Query().Get("uid")

	if uid == "" {
		// if uid is empty return an error
		fmt.Fprint(w, sendJsonGet("null", "uid not present in querystring"))
	} else {
		output, err := get(uid)
		if err != nil {
			fmt.Fprint(w, sendJsonGet("null", err.Error()))
		} else {
			// Check to see if the value is not set
			if output.Empty {
				// Return default if not set
				fmt.Fprint(w, sendJsonGet("#181A1B", "No value found: Default returned"))
			} else {
				fmt.Fprint(w, sendJsonGet(output.BG, "null"))
			}
		}
	}

}

func getTC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	uid := r.URL.Query().Get("uid")

	if uid == "" {
		// if uid is empty return an error
		fmt.Fprint(w, sendJsonGet("null", "uid not present in querystring"))
	} else {
		output, err := get(uid)
		if err != nil {
			fmt.Fprint(w, sendJsonGet("null", err.Error()))
		} else {
			// Check to see if the value is not set
			if output.Empty {
				// Return default if not set
				fmt.Fprint(w, sendJsonGet("#FFFFFF", "No value found: Default returned"))
			} else {
				fmt.Fprint(w, sendJsonGet(output.TC, "null"))
			}
		}
	}

}

func getSTC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	uid := r.URL.Query().Get("uid")

	if uid == "" {
		// if uid is empty return an error
		fmt.Fprint(w, sendJsonGet("null", "uid not present in querystring"))
	} else {
		output, err := get(uid)
		if err != nil {
			fmt.Fprint(w, sendJsonGet("null", err.Error()))
		} else {
			// Check to see if the value is not set
			if output.Empty {
				// Return default if not set
				fmt.Fprint(w, sendJsonGet("#D3D3D3", "No value found: Default returned"))
			} else {
				fmt.Fprint(w, sendJsonGet(output.STC, "null"))
			}
		}
	}
}
