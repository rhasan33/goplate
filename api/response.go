package api

import (
	"encoding/json"
	"net/http"
)

type response map[string]interface{}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}, args ...string) {
	var resmsg = response{}

	if payload != nil {
		if len(args) == 0 {
			resmsg["data"] = payload
		} else {
			typ := args[0]
			resmsg[typ] = []interface{}{payload}
		}
	}

	res, _ := json.Marshal(resmsg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg}, "errors")
}

// respondWithValidationError return error message
func respondWithValidationError(w http.ResponseWriter, code int, i interface{}) {
	respondwithJSON(w, code, i, "errors")
}
