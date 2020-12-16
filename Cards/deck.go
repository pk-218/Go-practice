package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// creating a new type of deck which stores types of cards as strings
type deck []string

// receiver function for printing cards in a deck type
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// creating a new deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// creating a hand of cards from the deck and returning both the hand and the remaining cards
func deal(d deck, handSize int) (deck, deck) {
	return d[handSize:], d[:handSize]
}

// convert deck which is a slice of strings to a single comma separated string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// receiver function to save a file to disk
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0600)
}

// read comma separated file to get new deck
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// receiver function to shuffle a deck
func (d deck) shuffle() {
	// creating a random number generator according to the current time so that
	// every time the program is run, a purely random number is generator
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i int, j int) {
		d[i], d[j] = d[j], d[i] // syntax for swapping d[i] and d[j]
	})
}
