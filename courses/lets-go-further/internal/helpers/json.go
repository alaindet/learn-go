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
	maxBytes *int64,
) error {
	var _maxBytes int64 = 2_097_152 // 2 Megabytes
	if maxBytes != nil {
		_maxBytes = *maxBytes
	}

	r.Body = http.MaxBytesReader(w, r.Body, _maxBytes)

	err := json.UnmarshalRead(
		r.Body,
		dest,
		json.RejectUnknownMembers(true),
		jsontext.AllowDuplicateNames(false),
	)

	if err == nil {
		return nil
	}

	var (
		semanticError  *json.SemanticError
		syntacticError *jsontext.SyntacticError
		maxBytesError  *http.MaxBytesError
	)

	switch {
	case errors.Is(err, io.EOF):
		return ErrEmptyJSON

	case errors.Is(err, io.ErrUnexpectedEOF):
		return ErrMalformedJSON

	case errors.As(err, &syntacticError):
		return fmt.Errorf(
			"body contains badly-formed JSON (at character %d)",
			syntacticError.ByteOffset,
		)

	case errors.As(err, &semanticError):
		return fmt.Errorf(
			"body contains badly-formed JSON (at character %d)",
			semanticError.ByteOffset,
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
