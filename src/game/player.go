package game

import (
	"fmt"

	"github.com/anwesh-b/go-tictactoe/src/player"
)

func togglePlayer(currentPlayer **player.Player, player1 *player.Player, player2 *player.Player) {
	if *currentPlayer == player1 {
		*currentPlayer = player2
	} else {
		*currentPlayer = player1
	}
}

func validatePlayerMove(gameState *Game, move player.Move) bool {
	for _, value := range &gameState.Moves {
		if value == nil {
			return true
		}

		if value.Position == move.Position {
			return false
		}
	}
	return false
}

func getPlayerInput(gameState *Game, currentPlayer *player.Player, moveIndex int8) {
	// Get input for the player.
	var pos uint8
	var newMove player.Move

	for isValidMove := false; isValidMove == false; {
		fmt.Printf("Hi %s, Please choose one cell(%s): ", currentPlayer.Name, player.CellRange)
		fmt.Scanln(&pos)

		var isValidSelection = player.IsMoveInRange(pos)
		if isValidSelection == false {
			fmt.Printf("Please choose cell between %s!!!\n\n", player.CellRange)
			continue
		}

		newMove = player.Move{
			Player:   currentPlayer,
			Position: player.Position(pos),
		}

		isValidMove = validatePlayerMove(gameState, newMove)

		if isValidMove == false {
			fmt.Printf("Cell already taken, please choose other one.\n\n")
		}
	}

	player.UpdatePlayerBoardMark(currentPlayer, pos)
	gameState.Moves[moveIndex] = &newMove
}
