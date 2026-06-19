package json

import "net/http"

type HTTPClient struct{}

func (client *HTTPClient) ReadJSON(
	w http.ResponseWriter,
	r *http.Request,
	toFill any,
) error {
	return ReadRequest(w, r, toFill)
}

func (client *HTTPClient) WriteJSON(
	w http.ResponseWriter,
	status int,
	jsonData any,
	headers ...http.Header,
) error {
	return WriteResponse(w, status, jsonData, headers...)
}

func (client *HTTPClient) WriteJSONError(
	w http.ResponseWriter,
	err error,
	status ...int,
) error {
	return WriteErrorResponse(w, err, status...)
}
