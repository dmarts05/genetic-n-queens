package population

import (
	"testing"
)

func Test_generateRandomIndividual(t *testing.T) {
	numQueens := 8
	ind := generateRandomIndividual(numQueens)

	// Check if the individual has no queens in the same row
	for i := 0; i < numQueens; i++ {
		for j := i + 1; j < numQueens; j++ {
			if ind.QueenPositions[i] == ind.QueenPositions[j] {
				t.Errorf("generateRandomIndividual() = %v, want no queens with the same row", ind)
			}
		}
	}
}

func TestGeneratePopulation(t *testing.T) {
	numQueens := 8
	populationSize := 100
	population := Generate(numQueens, populationSize)

	// Check if the population has the correct size
	if len(population) != populationSize {
		t.Errorf("GeneratePopulation() = %v, want %v", len(population), populationSize)
	}
}
