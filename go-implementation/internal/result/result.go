package result

import (
	"encoding/json"
	"fmt"
	"os"
)

// Represents the result of a single generation of the genetic algorithm
type GenerationResult struct {
	BestQueenPositions []int   `json:"best_queen_positions"`
	Generation         int     `json:"generation"`
	BestFitness        int     `json:"best_fitness"`
	MeanFitness        float64 `json:"mean_fitness"`
	IsSolution         bool    `json:"is_solution"`
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

// Get the best fitness of a slice of generation results
func GetBestFitness(results []GenerationResult) int {
	bestFitness := results[0].BestFitness
	for _, result := range results {
		if result.BestFitness > bestFitness {
			bestFitness = result.BestFitness
		}
	}
	return bestFitness
}

// Get the worst fitness of a slice of generation results
func GetWorstFitness(results []GenerationResult) int {
	worstFitness := results[0].BestFitness
	for _, result := range results {
		if result.BestFitness < worstFitness {
			worstFitness = result.BestFitness
		}
	}
	return worstFitness
}

// Get the mean of the mean fitness of a slice of generation results
func GetMeanMeanFitness(results []GenerationResult) float64 {
	totalFitness := 0.0
	for _, result := range results {
		totalFitness += result.MeanFitness
	}
	return totalFitness / float64(len(results))
}

// Get the mean best fitness of a slice of generation results
func GetMeanBestFitness(results []GenerationResult) float64 {
	totalFitness := 0.0
	for _, result := range results {
		totalFitness += float64(result.BestFitness)
	}
	return totalFitness / float64(len(results))
}

// Get the mean generations required to find the solution of a slice of generation results
func GetMeanGenerations(results []GenerationResult) float64 {
	totalGenerations := 0.0
	for _, result := range results {
		totalGenerations += float64(result.Generation)
	}
	return totalGenerations / float64(len(results))
}

func GetNumSolutions(results []GenerationResult) int {
	numSolutions := 0
	for _, result := range results {
		if result.IsSolution {
			numSolutions++
		}
	}
	return numSolutions
}
