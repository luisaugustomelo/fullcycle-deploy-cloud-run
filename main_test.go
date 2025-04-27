package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidZip(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/weather?cep=123", nil)
	w := httptest.NewRecorder()

	weatherHandler(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status 422, got %d", w.Code)
	}
}

func TestEmptyZip(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/weather", nil)
	w := httptest.NewRecorder()

	weatherHandler(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status 422, got %d", w.Code)
	}
}
