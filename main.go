package main

import (
	"service"
	"gateTimer"
)

func main() {
	go gateTimer.NewGateTimer()
	service.Init()
}