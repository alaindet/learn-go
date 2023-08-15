package common

import (
	"fmt"
	"greenlight/cmd/api/core"
)

func BackgroundTask(app *core.Application, fn func()) {
	go func() {

		// Panic recovery logic
		defer func(app *core.Application) {
			err := recover()
			if err != nil {
				app.Logger.Error(fmt.Errorf("%s", err).Error(), nil)
			}
		}(app)

		fn()

	}()
}
