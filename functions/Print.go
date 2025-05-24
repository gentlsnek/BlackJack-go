package functions

import (
	"fmt"
	"strings"
)

// Add this helper function to your main.go or a utils file
func prettyCard(card string) string {
	rankMap := map[string]string{
		"A": "Ace", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7",
		"8": "8", "9": "9", "10": "10", "J": "Jack", "Q": "Queen", "K": "King",
	}
	suitMap := map[string]string{
		"H": "Hearts", "S": "Spades", "C": "Clubs", "D": "Diamonds",
	}
	parts := strings.Split(card, "_")
	if len(parts) != 2 {
		return card
	}
	rank, suit := parts[0], parts[1]
	return fmt.Sprintf("%s of %s", rankMap[rank], suitMap[suit])
}

// Use this to print hands:
func Print(hand []string) []string {
	var output []string
	for _, card := range hand {
		output = append(output, prettyCard(card))
	}

	return output
}
