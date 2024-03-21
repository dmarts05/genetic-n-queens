package individual

import (
	"errors"
	"math/rand/v2"
	"slices"
)

// Represents an individual in the population
// QueenPositions: The positions of the queens on the board. Each index in the array represents the column of the queen and the value at that index represents the row of the queen
type Individual struct {
	QueenPositions []int
}

// Calculate the number of clashes between the queens for the individual
func (ind Individual) getNumClashes() int {
	numQueens := len(ind.QueenPositions)
	clashes := 0

	// Since every queen is in a different column and row, we only need to check for diagonal attacks
	// If the difference between the rows and columns of the 2 queens are equal, they are attacking each other
	for col1 := 0; col1 < numQueens; col1++ {
		for col2 := col1 + 1; col2 < numQueens; col2++ {
			row1 := ind.QueenPositions[col1]
			row2 := ind.QueenPositions[col2]

			if row1-col1 == row2-col2 || row1+col1 == row2+col2 {
				clashes++
			}
		}
	}

	return clashes
}

// Calculate the fitness of the individual
func (ind Individual) Fitness() int {
	numQueens := len(ind.QueenPositions)
	maxNonAttackingPairs := numQueens * (numQueens - 1) / 2
	clashes := ind.getNumClashes()
	fitness := maxNonAttackingPairs - clashes
	return fitness
}

// Perform crossover between two individuals to create two new individuals
// Here we are using OX because it let us avoid creating invalid individuals (i.e. individuals with duplicate queen positions or in the same row or column
func (ind Individual) Crossover(other Individual) (Individual, Individual, error) {
	// Check if the two individuals have the same amount of queens
	if len(ind.QueenPositions) != len(other.QueenPositions) {
		return Individual{}, Individual{}, errors.New("individuals have different number of queens")
	}

	// Create two new individuals to store the children
	numQueens := len(ind.QueenPositions)
	child1 := Individual{QueenPositions: make([]int, numQueens)}
	child2 := Individual{QueenPositions: make([]int, numQueens)}
	// Set the queen positions of the children to -1 to indicate that they are empty
	for i := 0; i < numQueens; i++ {
		child1.QueenPositions[i] = -1
		child2.QueenPositions[i] = -1
	}

	// Select two random points to perform the crossover
	point1 := rand.IntN(numQueens)
	point2 := rand.IntN(numQueens - 1)
	if point2 >= point1 {
		point2++
	} else {
		point1, point2 = point2, point1
	}

	// Copy the selected part of the parents to the children
	for i := point1; i < point2; i++ {
		child1.QueenPositions[i] = other.QueenPositions[i]
		child2.QueenPositions[i] = ind.QueenPositions[i]
	}

	// Repair the children by adding the missing queens

	// Child 1
	for i := 0; i < len(child1.QueenPositions); i++ {
		if i < point1 || i >= point2 && child1.QueenPositions[i] == -1 {
			for j := 0; j < len(ind.QueenPositions); j++ {
				if !slices.Contains(child1.QueenPositions, ind.QueenPositions[j]) {
					child1.QueenPositions[i] = ind.QueenPositions[j]
					break
				}
			}
		}
	}

	// Child 2
	for i := 0; i < len(child2.QueenPositions); i++ {
		if i < point1 || i >= point2 && child2.QueenPositions[i] == -1 {
			for j := 0; j < len(other.QueenPositions); j++ {
				if !slices.Contains(child2.QueenPositions, other.QueenPositions[j]) {
					child2.QueenPositions[i] = other.QueenPositions[j]
					break
				}
			}
		}
	}

	return child1, child2, nil
}

// Mutate the individual by shuffling the queen positions
func (ind *Individual) Mutate() {
	rand.Shuffle(len(ind.QueenPositions), func(i, j int) {
		ind.QueenPositions[i], ind.QueenPositions[j] = ind.QueenPositions[j], ind.QueenPositions[i]
	})
}
