package selection

import (
	"reflect"
	"testing"

	"github.com/dmarts05/genetic-n-queens/internal/individual"
)

func TestSelectByElitism(t *testing.T) {
	perfectIndividual := &individual.Individual{QueenPositions: []int{0, 4, 7, 5, 2, 6, 1, 3}}
	badIndividual := &individual.Individual{QueenPositions: []int{0, 1, 2, 3, 4, 5, 6, 7}}

	population := []*individual.Individual{
		perfectIndividual,
		badIndividual,
		badIndividual,
		badIndividual,
		badIndividual,
		badIndividual,
		badIndividual,
		perfectIndividual,
		badIndividual,
		badIndividual,
	}

	selected := SelectByElitism(population, 2)
	if len(selected) != 2 {
		t.Errorf("Expected 2 individuals, got %d", len(selected))
	}
	if !reflect.DeepEqual(selected[0], perfectIndividual) {
		t.Errorf("Expected the first individual to be the perfect individual")
	}
	if !reflect.DeepEqual(selected[1], perfectIndividual) {
		t.Errorf("Expected the second individual to be the perfect individual")
	}
}
