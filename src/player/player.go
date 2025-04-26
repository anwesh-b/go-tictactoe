package player

import "fmt"

type PlayerStyle string

const (
	X PlayerStyle = "x"
	O PlayerStyle = "o"
)

type Player struct {
	Name       string
	Style      PlayerStyle
	Id         string
	BoardMarks uint16
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
	Player   *Player
	Position Position
}

var minPos uint8 = 1
var maxPos uint8 = 9
var CellRange = fmt.Sprintf("%d-%d", minPos, maxPos)

func UpdatePlayerBoardMark(player *Player, newBoardPosition uint8) {
	var currentPlayerMark uint16

	// Reducing 1 to since bitwise shift starts from 0 position.
	currentPlayerMark |= 1 << (newBoardPosition - 1)

	var newPlayerMask = currentPlayerMark | player.BoardMarks

	player.BoardMarks = newPlayerMask
}

func IsMoveInRange(movePosition uint8) bool {
	return movePosition <= maxPos && movePosition >= minPos
}
