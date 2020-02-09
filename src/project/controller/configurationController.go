package controller

import (
	"io/ioutil"
	"log"
	"net/http"
)

func LoadConfiguration() []byte {
	resp, err := http.Get("https://raw.githubusercontent.com/tim-klug/gateTimerConfig/master/config.json")
	if err != nil {
		log.Println(err)
		return nil
	}
	if resp != nil {
		defer func() {
			if err := resp.Body.Close(); err != nil {
				log.Println(err)
			}
		}()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	return body
}
