package functions

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Load() {
	filePath := "./game_history.csv"
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Could not open game history:", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for _, rec := range records {
		if len(rec) == 4 {
			fmt.Printf("Game %s | Dealer: %s | Player: %s | Outcome: %s\n", rec[0], rec[1], rec[2], rec[3])
		} else {
			fmt.Println(rec)
		}
	}
}
