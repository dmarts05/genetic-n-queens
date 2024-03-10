package individual

import (
	"errors"
	"math"
	"math/rand/v2"
)

// Represents a position on the board
type Position struct {
	row    int
	column int
}

// Generate a random position on the board
func generateRandomPosition(boardSize int) Position {
	return Position{
		row:    rand.IntN(boardSize),
		column: rand.IntN(boardSize),
	}
}

// Check if 2 queens are attacking each other
func areQueensAttacking(pos1 Position, pos2 Position) bool {
	// Row and column check
	if pos1.row == pos2.row || pos1.column == pos2.column {
		return true
	}

	// Diagonal check
	rowDiff := int(math.Abs(float64(pos1.row - pos2.row)))
	colDiff := int(math.Abs(float64(pos1.column - pos2.column)))
	return rowDiff == colDiff || rowDiff == 0 || colDiff == 0
}

// Represents an individual in the population
// The board is represented as a 2D slice of booleans
// where true represents a queen and false represents an empty cell
type Individual struct {
	Board [][]bool
}

// Get the positions of all the queens on the board
func (ind *Individual) getQueenPositions() []Position {
	numQueens := len(ind.Board)
	queenPositions := make([]Position, numQueens)
	currentPositionIndex := 0
	for i, row := range ind.Board {
		for j, isQueen := range row {
			if isQueen {
				queenPositions[currentPositionIndex] = Position{row: i, column: j}
				currentPositionIndex++
			}
		}
	}
	return queenPositions
}

// Calculate the fitness of the individual
func (ind *Individual) Fitness() int {
	// Get all queen positions
	queenPositions := ind.getQueenPositions()

	// Calculate the number of clashes between queens
	clashes := 0
	for i := 0; i < len(queenPositions); i++ {
		for j := i + 1; j < len(queenPositions); j++ {
			if areQueensAttacking(queenPositions[i], queenPositions[j]) {
				clashes++
			}
		}
	}

	maxNonAttackingPairs := (len(ind.Board) * (len(ind.Board) - 1)) / 2
	fitness := maxNonAttackingPairs - clashes
	return fitness
}

// Perform crossover between two individuals to create two new individuals
func (ind *Individual) Crossover(other *Individual) (*Individual, *Individual, error) {
	// Check if the two individuals have the same board size
	if len(ind.Board) != len(other.Board) {
		return nil, nil, errors.New("individuals have different board sizes")
	}
	boardSize := len(ind.Board)

	// Randomly select a crossover point
	crossoverPoint := rand.IntN(len(ind.Board))

	// Create the children
	child1 := &Individual{Board: make([][]bool, len(ind.Board))}
	child2 := &Individual{Board: make([][]bool, len(ind.Board))}
	for i := 0; i < len(ind.Board); i++ {
		child1.Board[i] = make([]bool, boardSize)
		child2.Board[i] = make([]bool, boardSize)
	}

	// Copy the first part of the board from the first parent to the first child
	for i := 0; i < crossoverPoint; i++ {
		for j := 0; j < boardSize; j++ {
			child1.Board[i][j] = ind.Board[i][j]
			child2.Board[i][j] = other.Board[i][j]
		}
	}

	// Copy the second part of the board from the second parent to the first child
	for i := crossoverPoint; i < len(ind.Board); i++ {
		for j := 0; j < boardSize; j++ {
			child1.Board[i][j] = other.Board[i][j]
			child2.Board[i][j] = ind.Board[i][j]
		}
	}

	return child1, child2, nil
}

// Mutate the individual
func (ind *Individual) Mutate() {
	// Get all queen positions
	queenPositions := ind.getQueenPositions()

	// Randomly select a queen to move
	numQueens := len(ind.Board)
	queenToMovePosition := queenPositions[rand.IntN(numQueens)]

	// Randomly select a new position for the queen that is not the current position
	// We also want to avoid positions that are already occupied by other queens
	newQueenPosition := generateRandomPosition(len(ind.Board))
	for newQueenPosition == queenToMovePosition || ind.Board[newQueenPosition.row][newQueenPosition.column] {
		newQueenPosition = generateRandomPosition(len(ind.Board))
	}

	// Move the queen
	ind.Board[queenToMovePosition.row][queenToMovePosition.column] = false
	ind.Board[newQueenPosition.row][newQueenPosition.column] = true
}
