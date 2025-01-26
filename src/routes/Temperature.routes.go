package routes

import(
	"github.com/gin-gonic/gin"

	"github.com/lalo64/go-kafka/src/controllers"
	"github.com/lalo64/go-kafka/src/service"
)


func TemperatureRoutes(router *gin.RouterGroup, temperatureService *service.TemperatureService) {
	temperatureController := controllers.NewTemperatureController(temperatureService)

	router.POST("/v1/temperature", temperatureController.CreateTemperature)
}