package population

import (
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
func GeneratePopulation(numQueens, populationSize int) []*individual.Individual {
	population := make([]*individual.Individual, populationSize)

	for i := 0; i < populationSize; i++ {
		population[i] = generateRandomIndividual(numQueens)
	}

	return population
}
