package functions

import (
	"math/rand"
	"time"
)

func Shuffle(deck []string) {

	rand.Seed(time.Now().UnixNano())
	// Shuffle the deck of cards
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)               // random index from 0 to i
		deck[i], deck[j] = deck[j], deck[i] // swap elements
	}
}
