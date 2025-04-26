package game

import (
	"fmt"

	"github.com/anwesh-b/go-tictactoe/src/player"
)

type Game struct {
	Player1 player.Player
	Player2 player.Player
	Moves   [9]*player.Move
}

func StartGame(player1 player.Player, player2 player.Player) {

	player1.Style = "x"
	player2.Style = "o"

	var game Game = Game{
		Player1: player1,
		Player2: player2,
	}

	var currentPlayer = &player1
	var moveIndex int8 = 0

	for {
		getPlayerInput(&game, currentPlayer, moveIndex)

		// Draw the board.
		drawBoard(&game)

		if game.IsGameWon(*currentPlayer) {
			fmt.Println("\n\nCongrats: ", currentPlayer.Name, " on winning the game.")
			break
		}

		if game.IsGameDraw() {
			fmt.Println("\n\nGame draw. GG")
			break
		}

		// Toggle player
		togglePlayer(&currentPlayer, &player1, &player2)
		moveIndex++
	}
}
