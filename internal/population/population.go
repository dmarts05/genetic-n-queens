package population

import (
	"errors"
	"math/rand/v2"

	"github.com/dmarts05/genetic-n-queens/internal/individual"
	"github.com/dmarts05/genetic-n-queens/internal/position"
)

// Generate a random individual
func generateRandomIndividual(numQueens int) *individual.Individual {
	queenPositions := make([]position.Position, numQueens)
	for i := 0; i < numQueens; i++ {
		// We only make the column random in order to avoid having queens in the same row, reducing the number of clashes by default
		queenPositions[i] = position.Position{
			Row:    i,
			Column: rand.IntN(numQueens),
		}
	}
	return individual.New(queenPositions)
}

// Generate a population of random individuals with the given number of queens and population size
func GeneratePopulation(numQueens, populationSize int) ([]*individual.Individual, error) {
	if numQueens < 4 {
		return nil, errors.New("it is not possible to solve the problem with less than 4 queens")
	} else if populationSize < 1 || populationSize%2 != 0 {
		return nil, errors.New("population size should be at least 2 and an even number")
	}

	population := make([]*individual.Individual, populationSize)

	for i := 0; i < populationSize; i++ {
		population[i] = generateRandomIndividual(numQueens)
	}

	return population, nil
}
