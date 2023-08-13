package handlers

import (
	"fmt"
	commonHttp "greenlight/cmd/api/common/http"
	"greenlight/cmd/api/core"
	"net/http"
	"time"
)

// GET /healthcheck
func HealthcheckHandler(app *core.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := commonHttp.JSONPayload{
			Message: fmt.Sprintf("healthcheck %s", time.Now()),
			Data: map[string]any{
				"status":      "available",
				"environment": app.Config.Env,
				"version":     app.Version,
				"char":        'a',
				"b64":         []byte("This is a string"),
			},
		}

		err := commonHttp.WriteJSON(w, http.StatusOK, data, nil)

		if err != nil {
			app.InternalServerErrorResponse(w, r, err)
		}
	}
}
