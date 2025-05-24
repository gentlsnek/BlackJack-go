package functions

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Save(playerHand []string, dealerHand []string, deck []string, outcome bool) {
	// Open the file in append mode, create if it doesn't exist
	filePath := "./game_history.txt"
	var gameNum int = 1

	// Check if file exists and count existing games
	if _, err := os.Stat(filePath); err == nil {
		f, err := os.Open(filePath)
		if err == nil {
			defer f.Close()
			reader := csv.NewReader(f)
			reader.FieldsPerRecord = -1
			records, _ := reader.ReadAll()
			for _, rec := range records {
				if len(rec) > 0 && strings.TrimSpace(rec[0]) != "" {
					if n, err := strconv.Atoi(rec[0]); err == nil && n >= gameNum {
						gameNum = n + 1
					}
				}
			}
		}
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening save file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, card := range playerHand {
		writer.Write([]string{strconv.Itoa(gameNum), "player", card})
	}
	for _, card := range dealerHand {
		writer.Write([]string{strconv.Itoa(gameNum), "dealer", card})
	}

	writer.Write([]string{strconv.Itoa(gameNum), "outcome", strconv.FormatBool(outcome)})

	fmt.Println("Game saved successfully!")
}
