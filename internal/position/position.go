package position

import "math/rand/v2"

// Represents a position on a board
type Position struct {
	Row    int
	Column int
}

// Generate a random position on a board with the given size
func GenerateRandomPosition(boardSize int) Position {
	return Position{
		Row:    rand.IntN(boardSize),
		Column: rand.IntN(boardSize),
	}
}
