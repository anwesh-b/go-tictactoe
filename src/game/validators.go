package game

import "github.com/anwesh-b/go-tictactoe/src/player"

// All winning possibilitites represented in bits
var winMasks = []uint16{
	// Rows
	0b111000000,
	0b000111000,
	0b000000111,
	// Column
	0b100100100,
	0b010010010,
	0b001001001,
	// Diagonal
	0b100010001,
	0b001010100,
}

func (game *Game) IsGameWon(player player.Player) bool {
	// Game cannot be won without at least 3 pieces drawn.
	// If starting player has 3 pieces, then the number of moves must be 5.
	// Checking against 5 moves and recusing 1 for array's index starting from 0.
	if (game.Moves[4]) == nil {
		return false
	}

	for _, mask := range winMasks {
		// Compare player's mark with winable masks using bitwise and.
		if (player.BoardMarks & mask) == mask {
			return true
		}
	}
	return false
}

func (game *Game) IsGameDraw() bool {
	// The max number of moves is 9 and if the game is not won, then the game is draw.
	if game.Moves[8] == nil {
		return false
	}

	return true
}
