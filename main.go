package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type PlayerStyle string

const (
	X PlayerStyle = "x"
	O PlayerStyle = "o"
)

type Player struct {
	name       string
	style      PlayerStyle
	id         string
	boardMarks uint16
}

type Position int

const (
	P1 Position = 1
	P2 Position = 2
	P3 Position = 3
	P4 Position = 4
	P5 Position = 5
	P6 Position = 6
	P7 Position = 7
	P8 Position = 8
	P9 Position = 9
)

type Move struct {
	player   *Player
	position Position
}

type Game struct {
	player1 Player
	player2 Player
	moves   [9]*Move
}

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

func (game *Game) isGameWon(player Player) bool {
	// Game cannot be won without at least 3 pieces drawn.
	// If starting player has 3 pieces, then the number of moves must be 5.
	// Checking against 5 moves and recusing 1 for array's index starting from 0.
	if (game.moves[4]) == nil {
		return false
	}

	for _, mask := range winMasks {
		// Compare player's mark with winable masks using bitwise and.
		if (player.boardMarks & mask) == mask {
			return true
		}
	}
	return false
}

func (game *Game) isGameDraw() bool {
	// The max number of moves is 9 and if the game is not won, then the game is draw.
	if game.moves[8] == nil {
		return false
	}

	return true
}

func drawBoard(game *Game) {
	// Initialize empty string for board marking.
	var boardMarks string = "         "
	for _, value := range game.moves {
		if value == nil {
			break
		}
		var postion = int(value.position) - 1
		boardMarks = boardMarks[:postion] + string(value.player.style) + boardMarks[postion+1:]
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

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateUuid() string {
	result := make([]byte, 12)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}
	return string(result)
}

func togglePlayer(currentPlayer **Player, player1 *Player, player2 *Player) {
	if *currentPlayer == player1 {
		*currentPlayer = player2
	} else {
		*currentPlayer = player1
	}
}

func validatePlayerMove(gameState *Game, move Move) bool {
	for _, value := range &gameState.moves {
		if value == nil {
			return true
		}

		if value.position == move.position {
			return false
		}
	}
	return false
}

func updatePlayerBoardMark(player *Player, newBoardPosition uint8) {
	var currentPlayerMark uint16

	// Reducing 1 to since bitwise shift starts from 0 position.
	currentPlayerMark |= 1 << (newBoardPosition - 1)

	var newPlayerMask = currentPlayerMark | player.boardMarks

	player.boardMarks = newPlayerMask
}

func getInput(gameState *Game, currentPlayer *Player, moveIndex int8) {
	// Get input for the player.
	var pos uint8
	var newMove Move

	for isValidMove := false; isValidMove != true; {
		fmt.Printf("Hi " + currentPlayer.name + " Please choose one cell(1-9): ")
		fmt.Scanln(&pos)

		if pos > 9 || pos < 1 {
			fmt.Println("Please choose cell between 1-9!!!")
			break
		}

		newMove = Move{
			player:   currentPlayer,
			position: Position(pos),
		}

		isValidMove = validatePlayerMove(gameState, newMove)

		if isValidMove == false {
			fmt.Println("Cell already taken, please choose other one.")
		}
	}

	updatePlayerBoardMark(currentPlayer, pos)
	gameState.moves[moveIndex] = &newMove
}

func startGame(player1 Player, player2 Player) {
	player1.style = "x"
	player2.style = "o"

	var game Game = Game{
		player1: player1,
		player2: player2,
	}

	var currentPlayer = &player1
	var moveIndex int8 = 0

	for {
		getInput(&game, currentPlayer, moveIndex)

		// Draw the board.
		drawBoard(&game)

		if game.isGameWon(*currentPlayer) {
			fmt.Println("\n\nCongrats: ", currentPlayer.name, " on winning the game.")
			break
		}

		if game.isGameDraw() {
			fmt.Println("\n\nGame draw. GG")
			break
		}

		// Toggle player
		togglePlayer(&currentPlayer, &player1, &player2)
		moveIndex++
	}
}

func main() {
	var player1 Player = Player{
		name: "Shyam",
		id:   generateUuid(),
	}
	var player2 Player = Player{
		name: "Ram",
		id:   generateUuid(),
	}

	fmt.Println("Starting game.....")
	startGame(player1, player2)
	fmt.Println("Ending game.....")
}
