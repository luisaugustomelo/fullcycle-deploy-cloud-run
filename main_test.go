package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock da função getCityByZip para testes
var originalGetCityByZip = getCityByZip
var originalGetTemperature = getTemperature

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

func TestZipNotFound(t *testing.T) {
	getCityByZip = func(zip string) (string, error) {
		return "", errCityNotFound
	}
	defer func() { getCityByZip = originalGetCityByZip }()

	req := httptest.NewRequest(http.MethodGet, "/weather?cep=12345678", nil)
	w := httptest.NewRecorder()

	weatherHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestWeatherAPIFailure(t *testing.T) {
	getCityByZip = func(zip string) (string, error) {
		return "São Paulo", nil
	}
	getTemperature = func(city string) (float64, error) {
		return 0, errWeatherAPIFailure
	}
	defer func() {
		getCityByZip = originalGetCityByZip
		getTemperature = originalGetTemperature
	}()

	req := httptest.NewRequest(http.MethodGet, "/weather?cep=12345678", nil)
	w := httptest.NewRecorder()

	weatherHandler(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestWeatherSuccess(t *testing.T) {
	getCityByZip = func(zip string) (string, error) {
		return "São Paulo", nil
	}
	getTemperature = func(city string) (float64, error) {
		return 25.0, nil
	}
	defer func() {
		getCityByZip = originalGetCityByZip
		getTemperature = originalGetTemperature
	}()

	req := httptest.NewRequest(http.MethodGet, "/weather?cep=29216090", nil)
	w := httptest.NewRecorder()

	weatherHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

var (
	errCityNotFound      = errorString("city not found")
	errWeatherAPIFailure = errorString("weather API failed")
)

type errorString string

func (e errorString) Error() string {
	return string(e)
}
