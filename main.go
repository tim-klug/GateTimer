package main

import (
	"project/gateTimer"
	"project/service"
)

func main() {
	go gateTimer.NewGateTimer()
	service.Init()
}
