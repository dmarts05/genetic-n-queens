package selection

import (
	"errors"
	"math/rand/v2"

	"github.com/dmarts05/genetic-n-queens/internal/individual"
)

func SelectByTournament(population []*individual.Individual) ([]*individual.Individual, error) {
	if len(population) < 2 {
		return nil, errors.New("population must have at least 2 individuals")
	}

	selected := []*individual.Individual{}
	for len(selected) < len(population) {
		// Get 2 random individuals
		i1 := rand.IntN(len(population))
		i2 := rand.IntN(len(population))
		// Ensure they are different
		for i1 == i2 {
			i2 = rand.IntN(len(population))
		}

		ind1 := population[i1]
		ind2 := population[i2]

		// Select the best one
		if ind1.Fitness() > ind2.Fitness() {
			selected = append(selected, ind1)
		} else {
			selected = append(selected, ind2)
		}
	}

	return selected, nil
}
