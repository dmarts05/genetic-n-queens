package selection

import (
	"reflect"
	"testing"

	"github.com/dmarts05/genetic-n-queens/internal/individual"
	"github.com/dmarts05/genetic-n-queens/internal/position"
)

func TestSelectByElitism(t *testing.T) {
	perfectIndividual := individual.New([]position.Position{
		{Row: 0, Column: 0},
		{Row: 1, Column: 6},
		{Row: 2, Column: 4},
		{Row: 3, Column: 7},
		{Row: 4, Column: 1},
		{Row: 5, Column: 3},
		{Row: 6, Column: 5},
		{Row: 7, Column: 2},
	})
	badIndividual := individual.New([]position.Position{
		{Row: 0, Column: 0},
		{Row: 1, Column: 1},
		{Row: 2, Column: 2},
		{Row: 3, Column: 3},
		{Row: 4, Column: 4},
		{Row: 5, Column: 5},
		{Row: 6, Column: 6},
		{Row: 7, Column: 7},
	})

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
