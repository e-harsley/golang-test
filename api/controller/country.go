package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-play/api/schemas"
	"github.com/go-play/api/services"
	"github.com/go-play/utils"
	"github.com/go-playground/validator"
)

type CountryController struct {
	service  services.CountryService
	validate *validator.Validate
}

func NewCountryController(service services.CountryService, validate validator.Validate) *CountryController {
	return &CountryController{
		service:  service,
		validate: &validate,
	}
}

func (cc *CountryController) AddCountry(ctx *gin.Context) {
	var request schemas.CountryRequestSchema

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.validate.StructCtx(ctx, request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	country, err := cc.service.AddCountry(request)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	ctx.JSON(http.StatusOK, country)
}

func (cc *CountryController) GetCountries(ctx *gin.Context) {
	countries, err := cc.service.GetCountries()

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	ctx.JSON(http.StatusOK, countries)

}

func (cc *CountryController) GetCountry(ctx *gin.Context) {
	id := ctx.Param("id")

	country, err := cc.service.GetCountry(id)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	ctx.JSON(http.StatusOK, country)
}

func (cc *CountryController) UpdateCountry(ctx *gin.Context) {
	var request schemas.CountryRequestSchema
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.validate.StructCtx(ctx, request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	country, err := cc.service.UpdateCountry(id, request)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	ctx.JSON(http.StatusOK, country)
}
