package gateTimer

import (
	"github.com/jasonlvhit/gocron"
	"fmt"
	"time"
	"encoding/json"
	"github.com/brian-armstrong/gpio"
)

const CONTROL_PIN = 11

var timerConfiguration = `[{"type":"daily","time":"10:00"}]`
var Events []Event
var controlPin = gpio.NewOutput(CONTROL_PIN, false)
var cron gocron.Scheduler

func NewGateTimer() {
	//ControlPin := gpio.NewOutput(CONTROL_PIN, false)
	controlPin.High()
	//controlPin = ControlPin
	defer controlPin.Close()
	setConfiguration(timerConfiguration)
}

func setConfiguration(configuration string) {
	jsonData := []byte(configuration)
	var loadedEvents []Event
	err := json.Unmarshal(jsonData, &loadedEvents)
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
		<- cron.Start()
	}
}

func openGate() {
	controlPin.Low()
	fmt.Println("Gate relais on")
	time.Sleep(5 * time.Second)
	controlPin.High()
	fmt.Println("Gate relais close")
}
