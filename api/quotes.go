package api

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type Quote struct {
	Quote  string `json:"text"`
	Author string `json:"author"`
	
}

func GetQuoteInJSON() Quote {
	content, err := ioutil.ReadFile("./data/quotes.json")
	if err != nil {
		return Quote{}
	}

	var quotes []Quote
	err = json.Unmarshal(content, &quotes)
	if err != nil {
		return Quote{}
	}

	if len(quotes) > 0 {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(1640) // Generate a random number between 0 and 1999
		return quotes[randomIndex]
	}

	return Quote{}
}
