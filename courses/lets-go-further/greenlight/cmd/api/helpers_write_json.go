package main

import (
	"encoding/json"
	"net/http"
)

type JSONPayload struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data"`
}

func (app *application) writeJSON(
	w http.ResponseWriter,
	httpStatus int,
	data JSONPayload,
	headers http.Header,
) error {

	message, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for headerName, headerValue := range headers {
		w.Header()[headerName] = headerValue
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(message)

	return nil
}
