package population

import (
	"testing"
)

func Test_generateRandomIndividual(t *testing.T) {
	numQueens := 8
	ind := generateRandomIndividual(numQueens)

	// Check if the individual has no queens with the same row
	for i := 0; i < numQueens; i++ {
		for j := i + 1; j < numQueens; j++ {
			if ind.QueenPositions[i].Row == ind.QueenPositions[j].Row {
				t.Errorf("generateRandomIndividual() = %v, want no queens with the same row", ind)
			}
		}
	}
}

func TestGeneratePopulation(t *testing.T) {
	numQueens := 8
	populationSize := 100
	population, _ := GeneratePopulation(numQueens, populationSize)

	// Check if the population has the correct size
	if len(population) != populationSize {
		t.Errorf("GeneratePopulation() = %v, want %v", len(population), populationSize)
	}

	_, err := GeneratePopulation(3, populationSize)
	if err == nil {
		t.Errorf("GeneratePopulation() = %v, want error", err)
	}

	_, err = GeneratePopulation(numQueens, 1)
	if err == nil {
		t.Errorf("GeneratePopulation() = %v, want error", err)
	}

	_, err = GeneratePopulation(numQueens, 3)
	if err == nil {
		t.Errorf("GeneratePopulation() = %v, want error", err)
	}
}
