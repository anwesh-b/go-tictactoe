package game

import "fmt"

func drawBoard(game *Game) {
	// Initialize empty string for board marking.
	var boardMarks string = "         "
	for _, value := range game.Moves {
		if value == nil {
			break
		}
		var postion = int(value.Position) - 1
		boardMarks = boardMarks[:postion] + string(value.Player.Style) + boardMarks[postion+1:]
	}

	fmt.Printf("\n")
	fmt.Printf("|\t|\t|\t|\n")
	fmt.Printf("|   %s\t|   %s\t|   %s\t|\n", string(boardMarks[0]), string(boardMarks[1]), string(boardMarks[2]))
	fmt.Printf("|_______|_______|_______|\n")
	fmt.Printf("|\t|\t|\t|\n")
	fmt.Printf("|   %s\t|   %s\t|   %s\t|\n", string(boardMarks[3]), string(boardMarks[4]), string(boardMarks[5]))
	fmt.Printf("|_______|_______|_______|\n")
	fmt.Printf("|   %s\t|   %s\t|   %s\t|\n", string(boardMarks[6]), string(boardMarks[7]), string(boardMarks[8]))
	fmt.Printf("|\t|\t|\t|\n\n")
}
