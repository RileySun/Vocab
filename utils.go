package main

import(
	"os"
	"log"
	"time"
	"math/rand"
	"image/color"
	"io/ioutil"
	"encoding/json"
)

var WHITE color.NRGBA = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func LoadAllData() []string {
	var data []string
	entries, _ := os.ReadDir("./Data")
	for _, e := range entries {
        data = append(data, e.Name())
    }
    return data
}

func LoadCardsFromJSON(filename string) []*Card {
	path := "./Data/" + filename
	
	data, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		log.Fatal(readErr)
	}
	
	var cards []*Card
	jsonErr := json.Unmarshal(data, &cards)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	
	return cards
}

func RandomizeQuiz(quiz *Quiz) *Quiz {
	newCards := quiz.Cards
	
	for i := range newCards {
    	j := rand.Intn(i + 1)
    	newCards[i], newCards[j] = newCards[j], newCards[i]
	}
	
	quiz.Cards = newCards
	
	return quiz
}