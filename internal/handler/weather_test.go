package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aronkst/go-cep-temperature/internal/handler"
	"github.com/aronkst/go-cep-temperature/internal/model"
)

type MockWeatherService struct {
	Temperature *model.Temperature
	Err         error
}

func (m *MockWeatherService) GetWeatherByCEP(cep string) (*model.Temperature, error) {
	return m.Temperature, m.Err
}

func TestGetWeatherByCEP_ValidCEP(t *testing.T) {
	mockService := &MockWeatherService{
		Temperature: &model.Temperature{Celsius: 30, Fahrenheit: 86, Kelvin: 303.15},
		Err:         nil,
	}

	handler := handler.NewWeatherHandler(mockService)

	req, err := http.NewRequest("GET", "/?cep=12345678", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler.GetWeatherByCEP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"temp_C":30,"temp_F":86,"temp_K":303.15}`
	if strings.Trim(responseRecorder.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), expected)
	}
}

func TestGetWeatherByCEP_InvalidCEP(t *testing.T) {
	mockService := &MockWeatherService{
		Temperature: nil,
		Err:         fmt.Errorf("invalid zipcode"),
	}

	handler := handler.NewWeatherHandler(mockService)

	req, err := http.NewRequest("GET", "/?cep=12345-678", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler.GetWeatherByCEP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}

	expected := "invalid zipcode"
	if strings.Trim(responseRecorder.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), expected)
	}
}

func TestGetWeatherByCEP_NotFound(t *testing.T) {
	mockService := &MockWeatherService{
		Temperature: nil,
		Err:         fmt.Errorf("can not find zipcode"),
	}

	handler := handler.NewWeatherHandler(mockService)

	req, err := http.NewRequest("GET", "/?cep=99999999", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler.GetWeatherByCEP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	expected := "can not find zipcode"
	if strings.Trim(responseRecorder.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), expected)
	}
}

func TestGetWeatherByCEP_InternalServerError(t *testing.T) {
	mockService := &MockWeatherService{
		Temperature: nil,
		Err:         fmt.Errorf("internal server error"),
	}

	handler := handler.NewWeatherHandler(mockService)

	req, err := http.NewRequest("GET", "/?cep=00000000", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler.GetWeatherByCEP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	expected := "internal server error"
	if strings.Trim(responseRecorder.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), expected)
	}
}
