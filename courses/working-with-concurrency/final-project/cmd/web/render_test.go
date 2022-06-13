package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfigAddDefaultData(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	ctx := getContext(r)
	r = r.WithContext(ctx)

	testApp.Session.Put(ctx, "flash", "flash")
	testApp.Session.Put(ctx, "warning", "warning")
	testApp.Session.Put(ctx, "error", "error")

	td := testApp.AddDefaultData(&TemplateData{}, r)

	if td.Flash != "flash" {
		t.Error("failed to get flash data")
	}

	if td.Warning != "warning" {
		t.Error("failed to get warning data")
	}

	if td.Error != "error" {
		t.Error("failed to get error data")
	}
}

func TestConfigIsAuthenticated(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	ctx := getContext(r)
	r = r.WithContext(ctx)

	auth := testApp.IsAuthenticated(r)

	if auth {
		t.Error("returns true for IsAuthenticated() when it should be false")
	}

	testApp.Session.Put(ctx, "userID", 1)
	auth = testApp.IsAuthenticated(r)

	if !auth {
		t.Error("returns false for IsAuthenticated() when it should be true")
	}
}

func TestConfigRender(t *testing.T) {
	// Override global var
	templatesPath = "./templates"

	res := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getContext(req)
	req = req.WithContext(ctx)

	testApp.render(res, req, "home.page.gohtml", &TemplateData{})

	if res.Code != 200 {
		t.Error("failed to render page")
	}
}
