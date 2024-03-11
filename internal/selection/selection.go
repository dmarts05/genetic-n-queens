package selection

import (
	"github.com/dmarts05/genetic-n-queens/internal/individual"
	"github.com/dmarts05/genetic-n-queens/internal/util"
)

func SelectByTournament(population []*individual.Individual) []*individual.Individual {
	selected := []*individual.Individual{}
	for len(selected) < len(population) {
		// Get 2 random individuals
		samples := util.Sample(selected, 2)
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
