package repository_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aronkst/go-cep-temperature/internal/model"
	"github.com/aronkst/go-cep-temperature/internal/repository"
)

func TestWeatherByCoordinatesRepository_Success(t *testing.T) {
	t.Setenv("TEST", "true")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseBody := `{"current_weather":{"temperature":30.0}}`
		w.Write([]byte(responseBody))
	}))
	defer server.Close()

	repo := repository.NewWeatherByCoordinatesRepository(server.URL)

	coordinates := &model.Coordinates{
		Latitude:  "123",
		Longitude: "321",
	}

	temperature, err := repo.GetWeather(coordinates)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := &model.Weather{
		Temperature: 30.0,
	}

	if temperature.Temperature != expected.Temperature {
		t.Errorf("Latitude mismatch: expected %v, got %v", expected.Temperature, temperature.Temperature)
	}
}

func TestWeatherByCoordinatesRepository_ErrorHttp(t *testing.T) {
	t.Setenv("TEST", "true")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "", http.StatusMovedPermanently)
	}))
	defer server.Close()

	repo := repository.NewWeatherByCoordinatesRepository(server.URL)

	coordinates := &model.Coordinates{
		Latitude:  "123",
		Longitude: "321",
	}

	_, err := repo.GetWeather(coordinates)
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	expectedErrorMsg := "can not find zipcode"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("Error message does not match expected. \nExpected to contain: %s\nGot: %s", expectedErrorMsg, err.Error())
	}
}

func TestWeatherByCoordinatesRepository_NotStatusOK(t *testing.T) {
	t.Setenv("TEST", "true")

	statusServerError := http.StatusInternalServerError

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", statusServerError)
	}))
	defer server.Close()

	repo := repository.NewWeatherByCoordinatesRepository(server.URL)

	coordinates := &model.Coordinates{
		Latitude:  "123",
		Longitude: "321",
	}

	_, err := repo.GetWeather(coordinates)
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	expectedErrorMsg := "can not find zipcode"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("Error message does not match expected. \nExpected to contain: %s\nGot: %s", expectedErrorMsg, err.Error())
	}
}

func TestWeatherByCoordinatesRepository_ErrorJsonDecoder(t *testing.T) {
	t.Setenv("TEST", "true")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseBody := `error`
		w.Write([]byte(responseBody))
	}))
	defer server.Close()

	repo := repository.NewWeatherByCoordinatesRepository(server.URL)

	coordinates := &model.Coordinates{
		Latitude:  "123",
		Longitude: "321",
	}

	_, err := repo.GetWeather(coordinates)
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	expectedErrorMsg := "can not find zipcode"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("Error message does not match expected. \nExpected to contain: %s\nGot: %s", expectedErrorMsg, err.Error())
	}
}

func TestWeatherByCoordinatesRepository_JsonBlank(t *testing.T) {
	t.Setenv("TEST", "true")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseBody := `{"current_weather":{"temperature":0.0}}`
		w.Write([]byte(responseBody))
	}))
	defer server.Close()

	repo := repository.NewWeatherByCoordinatesRepository(server.URL)

	coordinates := &model.Coordinates{
		Latitude:  "123",
		Longitude: "321",
	}

	_, err := repo.GetWeather(coordinates)
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	expectedErrorMsg := "can not find zipcode"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("Error message does not match expected. \nExpected to contain: %s\nGot: %s", expectedErrorMsg, err.Error())
	}
}
