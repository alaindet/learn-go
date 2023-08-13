package http

import (
	"encoding/json"
	"net/http"
)

type JSONPayload struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"errors,omitempty"`
}

func WriteJSON(
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
