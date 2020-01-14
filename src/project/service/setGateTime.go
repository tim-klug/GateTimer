package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"project/gateTimer"
)

func NewSetGateTime(routerGroup gin.RouterGroup) {
	routerGroup.POST("/setGateTime", ResetGateTime)
	routerGroup.OPTIONS("/setGateTime", OptionsGate)
}

func ResetGateTime(context *gin.Context) {
	var configuration []gateTimer.Event
	err := context.BindJSON(&configuration)
	if err != nil {
		log.Fatalln(err)
	}
	if len(configuration) > 0 {
		gateTimer.Events = configuration
		go gateTimer.StartGateTimer()
		context.JSON(202, gateTimer.Events)
		return
	}
	context.JSON(400, gateTimer.Events)
}
