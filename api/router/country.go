package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-play/api/controller"
	"github.com/go-play/api/repository"
	"github.com/go-play/api/services"
	"github.com/go-playground/validator"
)

var countrycontroller = controller.NewCountryController(*services.NewCountryService(*repository.CountryRepo), *validator.New())

func CountryRouteInit(route *gin.RouterGroup) {
	rou := route.Group("/country")

	rou.POST("", countrycontroller.AddCountry)
	rou.PUT("/:id", countrycontroller.UpdateCountry)
	rou.GET("", countrycontroller.GetCountries)
	rou.GET("/:id", countrycontroller.GetCountry)
}
