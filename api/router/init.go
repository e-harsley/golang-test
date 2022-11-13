package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(route *gin.Engine) {
	countryRoute := route.Group("/api")

	CountryRouteInit(countryRoute)

	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "[Not found] route not found"})
	})
}
