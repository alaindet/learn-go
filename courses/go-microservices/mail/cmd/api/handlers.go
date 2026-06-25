package main

import (
	commonJSON "common/json"
	"fmt"
	"net/http"
)

type mailMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (app *App) SendMail(w http.ResponseWriter, r *http.Request) {
	var requestPayload mailMessage

	if err := app.ReadJSON(w, r, &requestPayload); err != nil {
		app.WriteJSONError(w, err)
		return
	}

	message := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	if err := app.Mailer.SendSMTPMessage(message); err != nil {
		app.WriteJSONError(w, err)
		return
	}

	app.WriteJSON(w, http.StatusAccepted, commonJSON.Response{
		Message: fmt.Sprintf("Email sent to %s", requestPayload.To),
	})
}
