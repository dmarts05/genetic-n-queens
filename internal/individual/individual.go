package individual

import (
	"errors"
	"math"
	"math/rand/v2"
	"slices"

	"github.com/dmarts05/genetic-n-queens/internal/position"
)

// Check if 2 queens are attacking each other
func areQueensAttacking(pos1, pos2 position.Position) bool {
	// Row and column check
	if pos1.Row == pos2.Row || pos1.Column == pos2.Column {
		return true
	}

	// Diagonal check
	rowDiff := int(math.Abs(float64(pos1.Row - pos2.Row)))
	colDiff := int(math.Abs(float64(pos1.Column - pos2.Column)))
	return rowDiff == colDiff || rowDiff == 0 || colDiff == 0
}

// Represents an individual in the population
type Individual struct {
	numQueens            int
	boardSize            int
	maxNonAttackingPairs int
	QueenPositions       []position.Position
}

func New(queenPositions []position.Position) *Individual {
	numQueens := len(queenPositions)
	boardSize := numQueens * numQueens
	maxNonAttackingPairs := (numQueens * (numQueens - 1)) / 2
	return &Individual{
		numQueens:            numQueens,
		boardSize:            boardSize,
		maxNonAttackingPairs: maxNonAttackingPairs,
		QueenPositions:       queenPositions,
	}
}

// Calculate the fitness of the individual
func (ind *Individual) Fitness() int {
	// Calculate the number of clashes between queens
	clashes := 0
	for i := 0; i < len(ind.QueenPositions); i++ {
		for j := i + 1; j < len(ind.QueenPositions); j++ {
			if areQueensAttacking(ind.QueenPositions[i], ind.QueenPositions[j]) {
				clashes++
			}
		}
	}

	fitness := ind.maxNonAttackingPairs - clashes
	return fitness
}

// Perform crossover between two individuals to create two new individuals
func (ind *Individual) Crossover(other *Individual) (*Individual, *Individual, error) {
	// Check if the two individuals have the same board size
	if ind.boardSize != other.boardSize {
		return nil, nil, errors.New("individuals have different board sizes")
	}

	// Randomly select a crossover point
	crossoverPoint := rand.IntN(len(ind.QueenPositions))

	// Create the children by slicing the queen positions of the parents
	child1 := New(append(ind.QueenPositions[:crossoverPoint], other.QueenPositions[crossoverPoint:]...))
	child2 := New(append(other.QueenPositions[:crossoverPoint], ind.QueenPositions[crossoverPoint:]...))

	return child1, child2, nil
}

// Mutate the individual
func (ind *Individual) Mutate() {
	// Randomly select a queen to move
	queenToMoveIndex := rand.IntN(ind.numQueens)

	// Randomly select a new position for the queen that is not the current position
	// We also want to avoid positions that are already occupied by other queens
	newQueenPosition := position.GenerateRandomPosition(ind.boardSize)
	for slices.Contains(ind.QueenPositions, newQueenPosition) {
		newQueenPosition = position.GenerateRandomPosition(ind.boardSize)
	}

	// Move the queen by updating the queen to move with the new position
	ind.QueenPositions = slices.Replace(ind.QueenPositions, queenToMoveIndex, queenToMoveIndex+1, newQueenPosition)
}
