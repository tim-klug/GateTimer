package main

import (
	"log"
	"project/gateTimer"
	"project/service"
)

func main() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("An error occurred: %s", err)
			}
		}()
		gateTimer.NewGateTimer()
	}()
	service.Init()
}
