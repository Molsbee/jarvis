package controller

import (
	"github.com/Molsbee/jarvis/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewCLCController() clcController {
	return clcController{}
}

type clcController struct {
}

func (c *clcController) FindIPAddress(context *gin.Context) {
	ipAddress := context.Param("ipAddress")
	address := service.FindIPAddress(ipAddress)
	if address.Address == ipAddress {
		context.JSON(http.StatusOK, address)
		return
	}

	context.Status(http.StatusNotFound)
}
