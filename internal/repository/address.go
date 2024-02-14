package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aronkst/go-cep-temperature/internal/model"
	"github.com/aronkst/go-cep-temperature/pkg/utils"
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
	defaultError := fmt.Errorf("invalid zipcode")

	if cep == "" || len(cep) != 8 || !utils.IsNumber(cep) {
		return nil, defaultError
	}

	var url string

	if os.Getenv("TEST") == "true" {
		url = repository.URL
	} else {
		url = fmt.Sprintf(repository.URL, cep)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, defaultError
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, defaultError
	}

	var address model.Address
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return nil, defaultError
	}

	if address.PostalCode == "" {
		return nil, defaultError
	}

	return &address, nil
}
