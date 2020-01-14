package controller

import (
	"io/ioutil"
	"log"
	"net/http"
)

func LoadConfiguration() []byte {
	resp, err := http.Get("https://raw.githubusercontent.com/tim-klug/gateTimerConfig/master/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
