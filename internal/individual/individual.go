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
	boardSize            int
	numQueens            int
	maxNonAttackingPairs int
	queenPositions       []position.Position
}

func New(queenPositions []position.Position) *Individual {
	boardSize := len(queenPositions)
	numQueens := len(queenPositions)
	maxNonAttackingPairs := (boardSize * (boardSize - 1)) / 2
	return &Individual{
		boardSize:            boardSize,
		numQueens:            numQueens,
		maxNonAttackingPairs: maxNonAttackingPairs,
		queenPositions:       queenPositions,
	}
}

// Calculate the fitness of the individual
func (ind *Individual) Fitness() int {
	// Calculate the number of clashes between queens
	clashes := 0
	for i := 0; i < len(ind.queenPositions); i++ {
		for j := i + 1; j < len(ind.queenPositions); j++ {
			if areQueensAttacking(ind.queenPositions[i], ind.queenPositions[j]) {
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
	crossoverPoint := rand.IntN(len(ind.queenPositions))

	// Create the children by slicing the queen positions of the parents
	child1 := New(append(ind.queenPositions[:crossoverPoint], other.queenPositions[crossoverPoint:]...))
	child2 := New(append(other.queenPositions[:crossoverPoint], ind.queenPositions[crossoverPoint:]...))

	return child1, child2, nil
}

// Mutate the individual
func (ind *Individual) Mutate() {
	// Randomly select a queen to move
	queenToMoveIndex := rand.IntN(ind.numQueens)

	// Randomly select a new position for the queen that is not the current position
	// We also want to avoid positions that are already occupied by other queens
	newQueenPosition := position.GenerateRandomPosition(ind.boardSize)
	for slices.Contains(ind.queenPositions, newQueenPosition) {
		newQueenPosition = position.GenerateRandomPosition(ind.boardSize)
	}

	// Move the queen by updating the queen to move with the new position
	ind.queenPositions = slices.Replace(ind.queenPositions, queenToMoveIndex, queenToMoveIndex+1, newQueenPosition)
}
