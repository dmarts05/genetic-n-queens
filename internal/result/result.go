package result

import (
	"encoding/json"

	"github.com/dmarts05/genetic-n-queens/internal/position"
)

// Represents the result of a single generation of the genetic algorithm
type GenerationResult struct {
	BestQueenPositions []position.Position `json:"best_queen_positions"`
	BestFitness        float64             `json:"best_fitness"`
	MeanFitness        float64             `json:"mean_fitness"`
}

// Transforms a GenerationResult into a JSON string
func (result GenerationResult) ToJSON() string {
	j, _ := json.MarshalIndent(result, "", "  ")
	return string(j)
}

// Get best generation result from a list of generation results
func GetBestGenerationResult(results []GenerationResult) GenerationResult {
	best := results[0]
	for _, result := range results {
		if result.BestFitness > best.BestFitness {
			best = result
		}
	}
	return best
}
