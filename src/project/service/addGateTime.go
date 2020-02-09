package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"project/gateTimer"
)

func NewAddGateTime(routerGroup gin.RouterGroup) {
	routerGroup.POST("/addGateTime", AddGateTime)
	routerGroup.OPTIONS("/addGateTime", OptionsGate)
}

func AddGateTime(context *gin.Context) {
	var configuration []gateTimer.Event
	err := context.BindJSON(&configuration)
	if err != nil {
		log.Println(err)
	}
	if len(configuration) > 0 {
		gateTimer.Events = append(gateTimer.Events, configuration...)
		go gateTimer.UpdateGateTimer()
		context.JSON(202, gateTimer.Events)
		return
	}
	context.JSON(400, gateTimer.Events)
}
