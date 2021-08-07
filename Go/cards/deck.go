package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { //_ is used for variables that we know exist but we don't care about
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { //reciever function, i.e. it will be called by a seq by .print()
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] //[:someInt] retrun beginning to someInt index
}

func (d deck) toString() string {
	//[]string (d) //type conversion of deck back to slice of string
	return strings.Join([]string(d), ",") // , is passed as the seperator
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //this func returns error //0666 is rw permission
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename) //returns byteslice and an error function if something went wrong
	if err != nil {                             //error occured
		fmt.Println("Error:", err)
		os.Exit(1) //close the entire programm entirely
	}
	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	//random seed generation - have to write random number ourselves
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
