package helpers

import (
	// "encoding/json/jsontext"
	"encoding/json/v2"
	"net/http"
)

func WriteJSON(
	w http.ResponseWriter,
	status int,
	data any,
	headers http.Header,
) error {
	// Prod
	jsonData, err := json.Marshal(data)

	// Dev
	// jsonData, err := json.Marshal(data, jsontext.Multiline(true))

	if err != nil {
		return err
	}

	jsonData = append(jsonData, '\n')

	for key, values := range headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
	return nil
}
