package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (app *application) readJSON(
	w http.ResponseWriter,
	r *http.Request,
	destination any,
) error {

	err := json.NewDecoder(r.Body).Decode(destination)

	// Everything's fine
	if err == nil {
		return nil
	}

	// Triaging the error...

	var errSyntax *json.SyntaxError
	var errUnmarshalType *json.UnmarshalTypeError
	var errInvalidUnmarshal *json.InvalidUnmarshalError

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

	// Empty
	case errors.Is(err, io.EOF):
		return errors.New("body must not be empty")

	// Generic
	case errors.As(err, &errInvalidUnmarshal):
		panic(err)
	default:
		return err
	}
}
