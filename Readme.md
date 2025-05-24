# Cards Game Project

This is a simple Go project for simulating card games, saving and loading game history in CSV format.

## Features

- Deal and shuffle cards
- Save each game’s player and dealer hands to a CSV file (`game_history.txt`)
- Incremental game numbering for each save
- Load and print game history from the CSV file

## Project Structure

```
cards/
├── functions/
│   ├── Deal.go
│   ├── Shuffle.go
│   ├── Save.go
│   ├── Load.go
│   └── (other utility files)
├── game_history.txt
├── main.go
└── Readme.md
```

## Usage

### Running the Project

1. **Initialize Go module (if not already):**
   ```sh
   go mod init cards
   go mod tidy
   ```

2. **Run the main program:**
   ```sh
   go run main.go
   ```

### Saving a Game

- The `Save` function in `functions/Save.go` appends the player and dealer hands to `game_history.txt` in CSV format.
- Each game is assigned an incrementing number in the first column.

### Loading and Printing Game History

- Use the `Load` function in `functions/Load.go` to read and print the game history.
- Use the `PrintGameHistory` function to print all records from the CSV file in a readable format.

## Example CSV Format

```
1,player,AS
1,player,10H
1,dealer,7C
1,dealer,5D
1,outcome,true
2,player,QS
2,player,8D
2,dealer,9C
2,dealer,6S
2,outcome,false
```

## Requirements

- Go 1.18 or newer

## License

MIT License