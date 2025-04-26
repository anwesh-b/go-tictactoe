package main

import (
	"fmt"

	"github.com/anwesh-b/go-tictactoe/src/game"
	"github.com/anwesh-b/go-tictactoe/src/player"
	"github.com/anwesh-b/go-tictactoe/src/utils"
)

func main() {
	var player1 player.Player = player.Player{
		Name: "Shyam",
		Id:   utils.GenerateUuid(),
	}
	var player2 player.Player = player.Player{
		Name: "Ram",
		Id:   utils.GenerateUuid(),
	}

	fmt.Println("Starting game.....")
	game.StartGame(player1, player2)
	fmt.Println("Ending game.....")
}
