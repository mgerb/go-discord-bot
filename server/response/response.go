package response

import (
	"encoding/json"
	"net/http"
)

var (
	DefaultUnauthorized  = []byte("Unauthorized.")
	DefaultInternalError = []byte("Internal error.")
)

// JSON - marshals the provided interface and returns JSON to client
func JSON(w http.ResponseWriter, content interface{}) {

	output, err := json.Marshal(content)

	if err != nil {
		ERR(w, http.StatusInternalServerError, []byte("Internal error."))
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

// ERR - send error response
func ERR(w http.ResponseWriter, status int, content []byte) {
	w.WriteHeader(status)
	w.Write(content)
}
