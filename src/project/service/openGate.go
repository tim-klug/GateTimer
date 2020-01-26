package service

import (
	"github.com/gin-gonic/gin"
	"project/gateTimer"
)

func NewOpenGate(routerGroup gin.RouterGroup) {
	routerGroup.GET("/openGate", OpenGateNow)
}

func OpenGateNow(context *gin.Context) {
	context.JSON(200, "OK")

	gateTimer.OpenGate()
}