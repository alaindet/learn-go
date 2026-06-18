package json

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

var (
	MaxJSONReadBytes = 1024 * 5
)

// Reads JSON payload (< 5 MB) an incoming HTTP request into a variable
func ReadRequest(
	w http.ResponseWriter,
	r *http.Request,
	data any,
) error {
	// Read the first 5 MB only
	r.Body = http.MaxBytesReader(w, r.Body, int64(MaxJSONReadBytes))

	// Read the entire body into memory
	dataBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// Unmarshal it into the "data" variable
	err = json.Unmarshal(dataBytes, data)
	if err != nil {
		return err
	}

	return nil
}

// Writes a JSON response
func WriteResponse(
	w http.ResponseWriter,
	status int,
	data any,
	headers ...http.Header,
) error {

	// Convert data to JSON
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Set HTTP headers response
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	// This is mandatory for JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Sending JSON bytes to the client
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func WriteErrorResponse(
	w http.ResponseWriter,
	err error,
	status ...int, // Here, status is a spread since it's optional and the last arg
) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload Response
	payload.Error = true
	payload.Message = err.Error()

	return WriteResponse(w, statusCode, payload)
}
