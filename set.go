package main

import (
	"fmt"
	"net/http"
)

func sendJsonSet(upsert int64, mod int64, err string) string {
	if err == "null" {
		json := fmt.Sprintf("{\"inserted\":%d,\"modified\":%d,\"err\":null}", upsert, mod)
		return json
	} else {
		json := fmt.Sprintf("{\"inserted\":%d,\"modified\":%d,\"err\":\"%v\"}", upsert, mod, err)
		return json
	}
}

func setConf(w http.ResponseWriter, r *http.Request) {
	// Check to see if bg, tc, stc, right, and uid are present in the querystring
	if r.URL.Query().Get("bg") == "" ||
		r.URL.Query().Get("tc") == "" ||
		r.URL.Query().Get("stc") == "" ||
		r.URL.Query().Get("right") == "" ||
		r.URL.Query().Get("uid") == "" {
		fmt.Fprint(w, "All inputs are required to use this endpoint")
	} else {
		var alignRight bool = false
		if r.URL.Query().Get("right") == "true" {
			alignRight = true
		}

		updateParams := user{
			BG:    r.URL.Query().Get("bg"),
			TC:    r.URL.Query().Get("tc"),
			STC:   r.URL.Query().Get("stc"),
			RIGHT: alignRight,
			UID:   r.URL.Query().Get("uid"),
		}

		result, err := set(updateParams)
		w.Header().Set("content-type", "application/json; charset=utf-8")
		if err != nil {
			fmt.Fprint(w, sendJsonSet(0, 0, err.Error()))
		} else {
			fmt.Fprint(w, sendJsonSet(result.UpsertedCount, result.ModifiedCount, "null"))
		}
	}

}
