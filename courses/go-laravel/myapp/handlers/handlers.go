package handlers

import (
	"net/http"

	"github.com/alaindet/gomitolo"
)

type Handlers struct {
	App *gomitolo.Gomitolo
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
