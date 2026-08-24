package helpers

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type JSONEnvelope map[string]any

func WriteJSON(
	w http.ResponseWriter,
	status int,
	data any,
	headers http.Header,
) error {
	jsonData, err := json.Marshal(data)
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

var (
	ErrEmptyJSON     = errors.New("body cannot be empty")
	ErrMalformedJSON = errors.New("body contains malformed JSON")
)

func ReadJSON[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var output T
	r.Body = http.MaxBytesReader(w, r.Body, 2_097_152)

	err := json.UnmarshalRead(
		r.Body,
		&output,
		json.RejectUnknownMembers(true),
		jsontext.AllowDuplicateNames(false),
	)

	if err == nil {
		return output, nil
	}

	var (
		syntacticError *jsontext.SyntacticError
		semanticError  *json.SemanticError
		maxBytesError  *http.MaxBytesError
	)

	switch {
	case errors.Is(err, io.EOF):
		return output, ErrEmptyJSON

	case errors.Is(err, io.ErrUnexpectedEOF):
		return output, ErrMalformedJSON

	case errors.As(err, &syntacticError):
		return output, fmt.Errorf(
			"body contains badly-formed JSON (at character %d)",
			syntacticError.ByteOffset,
		)

	case errors.As(err, &semanticError):
		return output, fmt.Errorf(
			"body contains invalid JSON (at character %d): %s",
			semanticError.ByteOffset,
			err.Error(),
		)

	case errors.As(err, &maxBytesError):
		return output, fmt.Errorf(
			"body must not be larger than %d bytes",
			maxBytesError.Limit,
		)

	default:
		return output, err
	}
}
