package main

import (
	"github.com/Molsbee/jarvis/server/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	clcController := controller.NewCLCController()

	router := gin.Default()
	router.GET("/clc/ipAddresses/:ipAddress", clcController.FindIPAddress)
	router.Run(":4000")
}
