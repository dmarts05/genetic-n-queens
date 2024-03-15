package result

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dmarts05/genetic-n-queens/internal/position"
)

// Represents the result of a single generation of the genetic algorithm
type GenerationResult struct {
	Generation         int                 `json:"generation"`
	BestQueenPositions []position.Position `json:"best_queen_positions"`
	BestFitness        int                 `json:"best_fitness"`
	MeanFitness        float64             `json:"mean_fitness"`
}

// Save a slice of generation results to a file in JSON format
func SaveResultsToFile(results []GenerationResult, path string) error {
	file, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		return fmt.Errorf("error marshalling results to JSON: %v", err)
	}

	err = os.WriteFile(path, file, 0644)
	if err != nil {
		return fmt.Errorf("error writing results to file: %v", err)
	}

	return nil
}
