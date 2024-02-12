package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aronkst/go-cep-temperature/internal/model"
)

type AddressRepository interface {
	GetEndereco(cep string) (*model.Address, error)
}

type addressRepository struct {
	URL string
}

func NewAddressRepository(url string) AddressRepository {
	return &addressRepository{
		URL: url,
	}
}

func (repository *addressRepository) GetEndereco(cep string) (*model.Address, error) {
	var url string

	if os.Getenv("TEST") == "true" {
		url = repository.URL
	} else {
		url = fmt.Sprintf(repository.URL, cep)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error when searching for zipcode %s information: %w", cep, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ViaCEP api returned status %d for zipcode %s: %w", resp.StatusCode, cep, err)
	}

	var address model.Address
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return nil, fmt.Errorf("error when decoding ViaCEP api response to zipcode %s: %w", cep, err)
	}

	return &address, nil
}
