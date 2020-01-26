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

func NewGateTimer() {

	fmt.Println("Checking functionality of the relay.")
	//ControlPin := gpio.NewOutput(CONTROL_PIN, false)
	err := controlPin.High()
	if err != nil {
		log.Fatalln(err)
	}
	//controlPin = ControlPin
	defer controlPin.Close()
	configuration := controller.LoadConfiguration()
	if configuration != nil {
		fmt.Println("Configuration was loaded %s", configuration)
		SetConfigurationByByte(configuration)
	} else {
		fmt.Println("The configuration could not be loaded, falling back to default config. Daily at 10:00.")
		standardEvent := Event{Type: "daily", Time: "10:00"}
		Events = append(Events, standardEvent)
	}

	UpdateGateTimer()
	createUpdateInterval()
}

func SetConfigurationByByte(configuration []byte) {
	var loadedEvents []Event
	err := json.Unmarshal(configuration, &loadedEvents)
	if err != nil {
		panic(err)
	}

	Events = loadedEvents
}

func UpdateGateTimer() {
	fmt.Println("Creating the scheduler for the events.")

	gocron.Clear()

	if len(Events) > 0 {
		for _, event := range Events {
			if event.Type == "daily" {
				fmt.Println("Add an event for the gate to open every day at " + event.Time)
				gocron.Every(1).Day().At(event.Time).Do(OpenGate)
			}
		}
		<-gocron.Start()
	}
}

func createUpdateInterval() {
	cron := gocron.NewScheduler()
	cron.Every(10).Minutes().Do(fetchConfiguration)
}

func fetchConfiguration() {
	configuration := controller.LoadConfiguration()
	if configuration != nil {
		SetConfigurationByByte(configuration)
	}
}

func OpenGate() {
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
