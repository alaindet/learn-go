package main

import (
	"greenlight/internal/validator"
	"net/url"
	"strconv"
	"strings"
)

func (app *application) readStringFromQueryString(
	queryString url.Values,
	key string,
	defaultValue string,
) string {
	s := queryString.Get(key)

	if s == "" {
		return defaultValue
	}

	return s
}

func (app *application) readCSVFromQueryString(
	queryString url.Values,
	key string,
	defaultValue []string,
) []string {

	csv := queryString.Get(key)

	if csv == "" {
		return defaultValue
	}

	return strings.Split(csv, ",")
}

func (app *application) readIntFromQueryString(
	queryString url.Values,
	key string,
	defaultValue int,
	v *validator.Validator,
) int {

	s := queryString.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)

	if err != nil {
		v.AddError(key, "must be an integer value")
		return defaultValue
	}

	return i
}