package functions

func Deal(deck []string, numcards int) ([]string, []string) {
	if numcards > len(deck) {
		numcards = len(deck)
	}
	hand := make([]string, numcards)
	copy(hand, deck[:numcards])
	deck = deck[numcards:]

	return hand, deck

}
