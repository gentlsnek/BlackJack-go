package main

import (
	"cards/functions"
	"fmt"
)

var card_deck []string = []string{"A_H", "2_H", "3_H", "4_H", "5_H", "6_H", "7_H", "8_H", "9_H", "10_H", "J_H", "Q_H", "K_H",
	"A_S", "2_S", "3_S", "4_S", "5_S", "6_S", "7_S", "8_S", "9_S", "10_S", "J_S", "Q_S", "K_S",
	"A_C", "2_C", "3_C", "4_C", "5_C", "6_C", "7_C", "8_C", "9_C", "10_C", "J_C", "Q_C", "K_C",
	"A_D", "2_D", "3_D", "4_D", "5_D", "6_D", "7_D", "8_D", "9_D", "10_D", "J_D", "Q_D", "K_D"}

func main() {
	player1pick := options1()

	switch player1pick {
	case 1:
		functions.Shuffle(card_deck[:])
		playGame()

	case 2:
		fmt.Println("Previous games")
		functions.Load()

	case 3:
		fmt.Println("Exiting")
		return

	default:
		fmt.Println("Invalid option, please try again.")
		return

	}
}
func options1() int32 {
	var options int32
	fmt.Println("1.New Games")
	fmt.Println("2.Previous Games")
	fmt.Println("3.Exit	")

	fmt.Scan(&options)

	return options
}

func playGame() {
	var playerHand []string
	var dealerHand []string
	var outcome string
	var gamestate bool
	var winloss bool

	playerHand, card_deck = functions.Deal(card_deck[:], 2)
	dealerHand, card_deck = functions.Deal(card_deck, 2)
	for {
		fmt.Println("Player Hand: ", playerHand)
		fmt.Println("Dealer Hand: ", dealerHand)

		fmt.Println("Pick an option:")
		fmt.Println("1. Hit")
		fmt.Println("2. Stand")
		var playerOption int32
		fmt.Scan(&playerOption)

		switch playerOption {
		case 1:
			var newCards []string
			newCards, card_deck = functions.Deal(card_deck, 1)
			playerHand = append(playerHand, newCards...)

		case 2:
			var newCards []string
			newCards, card_deck = functions.Deal(card_deck, 1)
			dealerHand = append(dealerHand, newCards...)

		default:
			fmt.Println("Invalid option, please try again.")
			continue
		}

		// Evaluate game status
		outcome, winloss, gamestate = functions.CheckWin(playerHand, dealerHand)

		if gamestate {
			fmt.Println(outcome)
			functions.Save(playerHand, dealerHand, card_deck, winloss)
			break
		}
	}
}
