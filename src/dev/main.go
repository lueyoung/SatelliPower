package main

import (
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Read config json error: ", err)
		return
	}
	log.Println("Read config json: ", string(data))
	var config *Config
	config = new(Config)
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatal("Json unmarshal error: ", err)
		return
	}
	foreach_struct(*config)
}
