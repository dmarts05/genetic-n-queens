package selection

import (
	"math/rand/v2"
	"sort"

	"github.com/dmarts05/genetic-n-queens/internal/individual"
	"github.com/dmarts05/genetic-n-queens/internal/util"
)

// Select individuals from the population using the tournament method
func SelectByTournament(population []*individual.Individual) []*individual.Individual {
	selected := []*individual.Individual{}
	for len(selected) < len(population) {
		// Get 2 random individuals
		samples := util.Sample(population, 2)
		ind1, ind2 := samples[0], samples[1]

		// Select the best one
		if ind1.Fitness() > ind2.Fitness() {
			selected = append(selected, ind1)
		} else {
			selected = append(selected, ind2)
		}
	}

	return selected
}

// Select individuals from the population using the roulette method
func SelectByRoulette(population []*individual.Individual) []*individual.Individual {
	selected := []*individual.Individual{}
	totalFitness := 0
	for _, ind := range population {
		totalFitness += ind.Fitness()
	}

	probabilities := make([]float64, len(population))
	for i, ind := range population {
		probabilities[i] = float64(ind.Fitness()) / float64(totalFitness)
	}

	cummulative_probabilities := make([]float64, len(population))
	cummulative_probabilities[0] = probabilities[0]
	for i := 1; i < len(population); i++ {
		cummulative_probabilities[i] = cummulative_probabilities[i-1] + probabilities[i]
	}

	for len(selected) < len(population) {
		r := rand.Float64()
		for i, p := range cummulative_probabilities {
			if r <= p {
				selected = append(selected, population[i])
				break
			}
		}
	}

	return selected
}

// Select n best individuals from the population
func SelectByElitism(population []*individual.Individual, n int) []*individual.Individual {
	sort.Slice(population, func(i, j int) bool {
		return population[i].Fitness() > population[j].Fitness()
	})
	return population[:n]
}
