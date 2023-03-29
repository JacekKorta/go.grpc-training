package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ParseJsonFile(path string) (Offers) {
	file, err := ioutil.ReadFile("../data/job_collection.json")
	if err != nil {
		log.Println(err)
	}

	var offers Offers
	err = json.Unmarshal(file, &offers)
	if err != nil {
		log.Println(err)
	}
	return offers
}