package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aronkst/go-cep-temperature/internal/service"
)

type WeatherHandler struct {
	weatherService service.WeatherService
}

func NewWeatherHandler(weatherService service.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		weatherService: weatherService,
	}
}

func (h *WeatherHandler) GetWeatherByCEP(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	temperature, err := h.weatherService.GetWeatherByCEP(cep)
	if err != nil {
		if err.Error() == "invalid zipcode" {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		} else {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(temperature)
}
