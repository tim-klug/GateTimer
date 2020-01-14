package service

import (
	"github.com/gin-gonic/gin"
	"project/gateTimer"
)

func NewShowGateTimes(routerGroup gin.RouterGroup) {
	routerGroup.GET("/gateTime", ShowGateTimes)
}

func ShowGateTimes(context *gin.Context) {
	context.JSON(200, gateTimer.Events)
}
