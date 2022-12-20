package main

import (
	"errors"
	"net/http"

	"github.com/go-playground/form/v4" // New import
)

func (app *application) decodePostForm(r *http.Request, target any) error {

	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(target, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError
		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
		return err
	}

	return nil
}
