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

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // anyone can read or write file
}

func newDeckFromFile(filename string) deck {
	byteSlice, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("error:", error)
		os.Exit(1)
	}

	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// func (d deck) shuffle() deck {
//   cardIndices := rand.Perm(52)
//   newDeck := deck{}
//   for i := range d {
//     newDeck = append(newDeck, d[cardIndices[i]])
//   }
//   return newDeck
// }

// func (d deck) shuffle() {
// 	rand.Seed(time.Now().UnixNano())
// 	rand.Shuffle(len(d), func(i, j int) {
// 		d[i], d[j] = d[j], d[i]
// 	})
// }
