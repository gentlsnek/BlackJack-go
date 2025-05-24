package functions

import (
	"strconv"
	"strings"
)

// CheckWin checks if the player has won the game
func CheckWin(playerHand []string, dealerHand []string) (string, bool) {
	playerScore := calculateScore(playerHand)
	dealerScore := calculateScore(dealerHand)
	if playerScore == 21 && len(playerHand) == 2 {
		return "Player wins with a Blackjack!  dealer score: " + strconv.Itoa(dealerScore) + " player score: " + strconv.Itoa(playerScore), true
	} else if dealerScore == 21 && len(dealerHand) == 2 {
		return "Dealer wins with a Blackjack!  dealer score: " + strconv.Itoa(dealerScore) + " player score: " + strconv.Itoa(playerScore), false
	} else if playerScore > 21 {
		return "Dealer wins! Player busted. dealer score: " + strconv.Itoa(dealerScore) + " player score: " + strconv.Itoa(playerScore), false
	} else if dealerScore > 21 {
		return "Player wins! Dealer busted. dealer score: " + strconv.Itoa(dealerScore) + " player score: " + strconv.Itoa(playerScore), true
	} else if playerScore > dealerScore {
		return "Player wins!  dealer score: " + strconv.Itoa(dealerScore) + " player score: " + strconv.Itoa(playerScore), true
	} else if dealerScore > playerScore {
		return "Dealer wins!  dealer score: " + strconv.Itoa(dealerScore) + " player score: " + strconv.Itoa(playerScore), false
	} else {
		return "It's a tie!  dealer score: " + strconv.Itoa(dealerScore) + " player score: " + strconv.Itoa(playerScore), false
	}
}

func calculateScore(hand []string) int {
	score := 0
	aces := 0

	for _, card := range hand {
		value := strings.Split(card, "_")[0]
		switch value {
		case "A":
			score += 11
			aces++
		case "K", "Q", "J":
			score += 10
		default:
			cardValue := strings.TrimLeft(value, "0")
			if cardValue == "" {
				cardValue = "10"
			}
			cardInt, _ := strconv.Atoi(cardValue)
			score += cardInt
		}
	}

	for score > 21 && aces > 0 {
		score -= 10
		aces--
	}

	return score
}
