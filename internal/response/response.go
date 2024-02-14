package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data any) {

	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := json.NewEncoder(w); err != nil {
		log.Fatal(err)
		return
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		ERROR string `json:"error"`
	}{ERROR: err.Error()})
}
