package service

import (
	"fmt"

	"github.com/aronkst/go-cep-temperature/internal/model"
	"github.com/aronkst/go-cep-temperature/internal/repository"
	"github.com/aronkst/go-cep-temperature/pkg/utils"
)

type WeatherService interface {
	GetWeatherByCEP(cep string) (*model.Temperature, error)
}

type weatherService struct {
	addressRepository              repository.AddressRepository
	coordinatesRepository          repository.CoordinatesRepository
	weatherByAddressRepository     repository.WeatherByAddressRepository
	weatherByCoordinatesRepository repository.WeatherByCoordinatesRepository
}

func NewWeatherService(
	addressRepository repository.AddressRepository,
	coordinatesRepository repository.CoordinatesRepository,
	weatherByAddressRepository repository.WeatherByAddressRepository,
	weatherByCoordinatesRepository repository.WeatherByCoordinatesRepository,
) WeatherService {
	return &weatherService{
		addressRepository:              addressRepository,
		coordinatesRepository:          coordinatesRepository,
		weatherByAddressRepository:     weatherByAddressRepository,
		weatherByCoordinatesRepository: weatherByCoordinatesRepository,
	}
}

func (s *weatherService) GetWeatherByCEP(cep string) (*model.Temperature, error) {
	address, err := s.addressRepository.GetEndereco(cep)
	if err != nil {
		return nil, fmt.Errorf("invalid zipcode")
	}

	var weather *model.Weather

	coordinates, err := s.coordinatesRepository.GetCoordinates(address)
	if err == nil {
		weather, err = s.weatherByCoordinatesRepository.GetWeather(coordinates)
		if err != nil {
			return nil, fmt.Errorf("can not find zipcode")
		}
	} else {
		weather, err = s.weatherByAddressRepository.GetWeather(address)
		if err != nil {
			return nil, fmt.Errorf("can not find zipcode")
		}
	}

	temperature := &model.Temperature{
		Celsius:    weather.Temperature,
		Fahrenheit: utils.CelsiusToFahrenheit(weather.Temperature),
		Kelvin:     utils.CelsiusToKelvin(weather.Temperature),
	}

	return temperature, nil
}
