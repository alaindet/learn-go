package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (app *application) readJSON(
	w http.ResponseWriter,
	r *http.Request,
	destination any,
) error {

	maxBytes := 1_048_576 // 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Setup JSON decoder
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	// Try decoding
	err := dec.Decode(destination)

	// Triaging the error
	if err != nil {

		var errSyntax *json.SyntaxError
		var errUnmarshalType *json.UnmarshalTypeError
		var errInvalidUnmarshal *json.InvalidUnmarshalError
		var errMaxBytes *http.MaxBytesError
		unknownFieldPrefix := "json: unknown field "

		switch {

		// Malformed
		case errors.As(err, &errSyntax):
			return fmt.Errorf("body contains malformed JSON (at char %d)", errSyntax.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains malformed JSON")

		// Wrong types
		case errors.As(err, &errUnmarshalType):
			field := errUnmarshalType.Field
			if field != "" {
				return fmt.Errorf("body contains wrong JSON type for field %q", field)
			}
			offset := errUnmarshalType.Offset
			return fmt.Errorf("body contains wrong JSON type at char %d", offset)

		// Unknown field
		case strings.HasPrefix(err.Error(), unknownFieldPrefix):
			fieldName := strings.TrimPrefix(err.Error(), unknownFieldPrefix)
			return fmt.Errorf("body contains unknown key %s", fieldName)

		// Empty
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		// Max size
		case errors.As(err, &errMaxBytes):
			return fmt.Errorf("body must not be larger than %d bytes", errMaxBytes.Limit)

		// Generic
		case errors.As(err, &errInvalidUnmarshal):
			panic(err)
		default:
			return err
		}
	}

	// Is there additional unexpected data in the body?
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain one valid JSON value")
	}

	// All's fine
	return nil
}
