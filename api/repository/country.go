package repository

import (
	"github.com/go-play/api/model"
	"github.com/go-play/api/schemas"
	"github.com/go-play/database"
	ormwrapper "github.com/go-play/orm_wrapper"
)

var CountryRepo = ormwrapper.NewRepository[schemas.CountryRequestSchema, model.Country](database.Init())
