package result

import (
	"reflect"
	"testing"

	"github.com/dmarts05/genetic-n-queens/internal/position"
)

func TestGenerationResult_ToJSON(t *testing.T) {
	result := GenerationResult{
		Generation: 1,
		BestQueenPositions: []position.Position{
			{Row: 0, Column: 0},
			{Row: 1, Column: 1},
			{Row: 2, Column: 2},
			{Row: 3, Column: 3},
		},
		BestFitness: 2,
		MeanFitness: 0.25,
	}
	expected := `{
  "generation": 1,
  "best_queen_positions": [
    {
      "row": 0,
      "column": 0
    },
    {
      "row": 1,
      "column": 1
    },
    {
      "row": 2,
      "column": 2
    },
    {
      "row": 3,
      "column": 3
    }
  ],
  "best_fitness": 2,
  "mean_fitness": 0.25
}`

	if result.ToJSON() != expected {
		t.Errorf("Expected %s, got %s", expected, result.ToJSON())
	}
}

func TestGetBestGenerationResult(t *testing.T) {
	results := []GenerationResult{
		{
			BestFitness: 0.5,
		},
		{
			BestFitness: 0.3,
		},
		{
			BestFitness: 0.7,
		},
	}
	expected := GenerationResult{
		BestFitness: 0.7,
	}

	if !reflect.DeepEqual(GetBestGenerationResult(results), expected) {
		t.Errorf("Expected %v, got %v", expected, GetBestGenerationResult(results))
	}
}
