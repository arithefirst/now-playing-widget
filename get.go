package main

import (
	"fmt"
	"net/http"
)

func getBG(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("uid")
	output, err := get(uid)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	} else {
		// Check to see if the value is not set
		if output.Empty {
			// Return default if not set
			fmt.Fprintf(w, "#181A1B")
		} else {
			fmt.Fprintf(w, output.BG)
		}
	}

}

func getTC(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("uid")
	output, err := get(uid)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	} else {
		// Check to see if the value is not set
		if output.Empty {
			// Return default if not set
			fmt.Fprintf(w, "#FFFFFF")
		} else {
			fmt.Fprintf(w, output.TC)
		}
	}

}

func getSTC(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("uid")
	output, err := get(uid)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	} else {
		// Check to see if the value is not set
		if output.Empty {
			// Return default if not set
			fmt.Fprintf(w, "#D3D3D3")
		} else {
			fmt.Fprintf(w, output.STC)
		}
	}
}
