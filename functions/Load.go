package functions

import (
	"fmt"
	"os"
	"strings"
)

func readfiles() {
	entries, err := os.ReadDir("./decks")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("Available saves:")
	for _, entry := range entries {
		fmt.Println("-", entry.Name())
	}

}

func Load() []string {
	readfiles()
	var filename string
	fmt.Println("Enter the name of the file to load:")
	fmt.Scan(&filename)

	file, err := os.ReadFile("./decks/" + filename)
	if err != nil {
		fmt.Println("Corrupt save file:", err)
		return nil
	}

	deck := string(file)
	deck = strings.TrimSpace(deck)     // Remove any leading/trailing whitespace
	cards := strings.Split(deck, "\n") // Split the string into a slice of cards

	return cards
}
