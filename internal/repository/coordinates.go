package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aronkst/go-cep-temperature/internal/model"
)

type CoordinatesRepository interface {
	GetCoordinates(address *model.Address) (*model.Coordinates, error)
}

type coordinatesRepository struct {
	URL        string
	BaseMethod string
}

func NewCoordinatesRepository(url string) CoordinatesRepository {
	return &coordinatesRepository{
		URL: url,
	}
}

func (repository *coordinatesRepository) GetCoordinates(address *model.Address) (*model.Coordinates, error) {
	defaultError := fmt.Errorf("can not find coordinates")

	baseURL := repository.URL

	params := url.Values{}
	params.Add("city", address.City)
	params.Add("state", address.State)
	params.Add("country", "Brasil")
	params.Add("format", "json")

	req, err := http.NewRequest("GET", baseURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, defaultError
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, defaultError
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, defaultError
	}

	var results []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, defaultError
	}

	if len(results) == 0 {
		return nil, defaultError
	}

	coordinates := &model.Coordinates{
		Latitude:  results[0].Lat,
		Longitude: results[0].Lon,
	}

	return coordinates, nil
}
