package functions

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Load() {
	filePath := "./game_history.txt"
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
		fmt.Println(rec)
	}
}
