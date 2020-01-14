package gateTimer

import (
	"encoding/json"
	"fmt"
	"github.com/brian-armstrong/gpio"
	"github.com/jasonlvhit/gocron"
	"log"
	"project/controller"
	"time"
)

const ControlPin = 11

var Events []Event
var controlPin = gpio.NewOutput(ControlPin, false)
var cron gocron.Scheduler

func NewGateTimer() {
	//ControlPin := gpio.NewOutput(CONTROL_PIN, false)
	err := controlPin.High()
	if err != nil {
		log.Fatalln(err)
	}
	//controlPin = ControlPin
	defer controlPin.Close()
	SetConfigurationByByte(controller.LoadConfiguration())
}

func SetConfigurationByByte(configuration []byte) {
	var loadedEvents []Event
	err := json.Unmarshal(configuration, &loadedEvents)
	if err != nil {
		panic(err)
	}

	Events = append(Events, loadedEvents...)

	StartGateTimer()
}

func StartGateTimer() {
	cron.Clear()
	cron := gocron.NewScheduler()

	if len(Events) > 0 {
		for _, event := range Events {
			if event.Type == "daily" {
				fmt.Println("Add an event for the gate to open every day at " + event.Time)
				cron.Every(1).Day().At(event.Time).Do(openGate)
			}
		}
		<-cron.Start()
	}
}

func openGate() {
	err := controlPin.Low()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Gate relays on")
	time.Sleep(5 * time.Second)
	err = controlPin.High()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Gate relays close")
}
