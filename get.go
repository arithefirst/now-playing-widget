package main

import (
	"fmt"
	"net/http"
)

// Function to create the json responses for the getX() funcs
func sendJson(hex string, err string) string {
	if hex == "null" {
		json := "{\"hex\":null,\"err\":\"" + err + "\"}"
		return json
	} else if err == "null" {
		json := "{\"hex\":\"" + hex + ",\"err\":null}"
		return json
	} else {
		json := "{\"hex\":\"" + hex + "\",\"err\":\"" + err + "\"}"
		return json
	}
}

func getBG(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	uid := r.Header.Get("uid")

	if uid == "" {
		// if uid is empty return an error
		fmt.Fprint(w, sendJson("null", "UID header empty"))
	} else {
		output, err := get(uid)
		if err != nil {
			fmt.Fprint(w, sendJson("null", err.Error()))
		} else {
			// Check to see if the value is not set
			if output.Empty {
				// Return default if not set
				fmt.Fprint(w, sendJson("#181A1B", "No value found: Default returned"))
			} else {
				fmt.Fprint(w, sendJson(output.BG, "null"))
			}
		}
	}

}

func getTC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	uid := r.Header.Get("uid")

	if uid == "" {
		// if uid is empty return an error
		fmt.Fprint(w, sendJson("null", "UID header empty"))
	} else {
		output, err := get(uid)
		if err != nil {
			fmt.Fprint(w, sendJson("null", err.Error()))
		} else {
			// Check to see if the value is not set
			if output.Empty {
				// Return default if not set
				fmt.Fprint(w, sendJson("#FFFFFF", "No value found: Default returned"))
			} else {
				fmt.Fprint(w, sendJson(output.TC, "null"))
			}
		}
	}

}

func getSTC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	uid := r.Header.Get("uid")

	if uid == "" {
		// if uid is empty return an error
		fmt.Fprint(w, sendJson("null", "UID header empty"))
	} else {
		output, err := get(uid)
		if err != nil {
			fmt.Fprint(w, sendJson("null", err.Error()))
		} else {
			// Check to see if the value is not set
			if output.Empty {
				// Return default if not set
				fmt.Fprint(w, sendJson("#D3D3D3", "No value found: Default returned"))
			} else {
				fmt.Fprint(w, sendJson(output.STC, "null"))
			}
		}
	}
}
