package schemas

import (
	"github.com/go-play/api/model"
)

type CountryRequestSchema struct {
	ID            uint   `json:"-"`
	Name          string `json:"name" validate:"required"`
	ShortName     string `json:"short_name" validate:"required"`
	Continent     string `json:"continent" validate:"required"`
	IsOperational *bool  `json:"is_operational"`
}

func (schema CountryRequestSchema) ToModel() model.Country {
	return model.Country{
		Name:          schema.Name,
		ShortName:     schema.ShortName,
		Continent:     schema.Continent,
		IsOperational: schema.IsOperational,
	}
}

func (schema CountryRequestSchema) FromModel(mc model.Country) interface{} {
	return CountryRequestSchema{
		ID:            mc.ID,
		Name:          mc.Name,
		ShortName:     mc.ShortName,
		Continent:     mc.Continent,
		IsOperational: mc.IsOperational,
	}
}
