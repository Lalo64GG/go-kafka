package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lalo64/go-kafka/src/infraestructure"
	"github.com/lalo64/go-kafka/src/models"
	"github.com/lalo64/go-kafka/src/service"
)


type TemperatureController struct {
	TemperatureService *service.TemperatureService
}

func NewTemperatureController(temperatureService *service.TemperatureService) *TemperatureController {
    return &TemperatureController{TemperatureService: temperatureService }
}



func (ctr *TemperatureController) CreateTemperature(ctx *gin.Context) {
	var temperature models.TemperatureModel
   
	if err := ctx.ShouldBind(&temperature); err != nil {
		ctx.JSON(http.StatusAccepted, models.Response{
			Success:  false,
            Message: "Error en los datos de la temperatura",
			Error: err.Error(),
			Data: nil,
		})
	}


	err := ctr.TemperatureService.CreateTemperature(temperature)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success:  false,
			Message: "Error al insertar la temperatura",
			Error: err.Error(),
			Data: nil,
		})
	}

	infraestructure.Producer(strconv.FormatFloat(temperature.Temperature, 'f', -1, 64))

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Registro de temperatura insertado correctamente",
		Error: "",
		Data: temperature,
	})
}