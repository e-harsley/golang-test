package services

import (
	"context"
	"strconv"

	"github.com/go-play/api/model"
	"github.com/go-play/api/schemas"
	ormwrapper "github.com/go-play/orm_wrapper"
)

var ctx = context.Background()

type CountryService struct {
	repository ormwrapper.ORMRepository[schemas.CountryRequestSchema, model.Country]
}

func NewCountryService(repo ormwrapper.ORMRepository[schemas.CountryRequestSchema, model.Country]) *CountryService {
	return &CountryService{
		repository: repo,
	}
}

func (cs *CountryService) AddCountry(data schemas.CountryRequestSchema) (*schemas.CountryRequestSchema, error) {
	err := cs.repository.Create(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (cs *CountryService) GetCountries() ([]schemas.CountryRequestSchema, error) {
	data, err := cs.repository.Find(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cs *CountryService) GetCountry(id string) (schemas.CountryRequestSchema, error) {
	uid, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return schemas.CountryRequestSchema{}, err
	}

	data, err := cs.repository.FindByID(ctx, uint(uid))
	if err != nil {
		return schemas.CountryRequestSchema{}, err
	}

	return data, nil
}

func (cs *CountryService) UpdateCountry(id string, data schemas.CountryRequestSchema) (*schemas.CountryRequestSchema, error) {
	uid, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return nil, err
	}

	err = cs.repository.Update(ctx, uint(uid), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil

}
