package helpers

import (
	// "encoding/json/jsontext"
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

var (
	ErrEmptyJSON     = errors.New("body cannot be empty")
	ErrMalformedJSON = errors.New("body contains malformed JSON")
)

func ReadJSON(
	w http.ResponseWriter,
	r *http.Request,
	dest any,
) error {
	// Set a hard limit of 2 MB
	// TODO: Make this configurable?
	r.Body = http.MaxBytesReader(w, r.Body, 2_097_152)

	err := json.UnmarshalRead(
		r.Body,
		dest,
		json.RejectUnknownMembers(true),
		jsontext.AllowDuplicateNames(false),
	)

	if err == nil {
		return nil
	}

	var semanticError *json.SemanticError
	var syntacticError *jsontext.SyntacticError
	var maxBytesError *http.MaxBytesError

	switch {
	case errors.Is(err, io.EOF):
		return ErrEmptyJSON

	case errors.Is(err, io.ErrUnexpectedEOF):
		return ErrMalformedJSON

	case errors.As(err, &semanticError):
		return fmt.Errorf(
			"semantic error at byte offset %d: %w",
			semanticError.ByteOffset,
			semanticError,
		)

	case errors.As(err, &syntacticError):
		return fmt.Errorf(
			"syntax error at byte offset %d: %w",
			semanticError.ByteOffset,
			semanticError,
		)

	case errors.As(err, &maxBytesError):
		return fmt.Errorf(
			"body must not be larger than %d bytes",
			maxBytesError.Limit,
		)

	default:
		return err
	}
}
