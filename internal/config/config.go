package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Represents the available selection methods for the genetic algorithm
type SelectionMethodType string

const (
	Tournament SelectionMethodType = "tournament"
	Roulette   SelectionMethodType = "roulette"
)

// Represents the configuration for the genetic algorithm
type Config struct {
	NumRuns         int                 `json:"num_runs"`
	SelectionMethod SelectionMethodType `json:"selection_method"`
	PopulationSize  int                 `json:"population_size"`
	MaxGenerations  int                 `json:"max_generations"`
	NumQueens       int                 `json:"num_queens"`
	MutationRate    float64             `json:"mutation_rate"`
	CrossOverRate   float64             `json:"crossover_rate"`
	Elitism         bool                `json:"elitism"`
}

// Load configuration from specified path or use default configuration if not found
func LoadConfig(path string) Config {
	defaultConfig := Config{
		NumRuns:         1,
		SelectionMethod: Roulette,
		PopulationSize:  16,
		MaxGenerations:  500,
		NumQueens:       8,
		MutationRate:    0.1,
		CrossOverRate:   0.5,
		Elitism:         true,
	}

	// Load config from specified path
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("[WARN] Specified configuration file not found, falling back to default configuration.")
		return defaultConfig
	}

	uncheckedConfig := struct {
		NumRuns         *int                 `json:"num_runs"`
		SelectionMethod *SelectionMethodType `json:"selection_method"`
		PopulationSize  *int                 `json:"population_size"`
		MaxGenerations  *int                 `json:"max_generations"`
		NumQueens       *int                 `json:"num_queens"`
		MutationRate    *float64             `json:"mutation_rate"`
		CrossOverRate   *float64             `json:"crossover_rate"`
		Elitism         *bool                `json:"elitism"`
	}{}

	// Load json into uncheckedConfig
	err = json.Unmarshal(data, &uncheckedConfig)
	if err != nil {
		fmt.Println("[WARN] Invalid config file, falling back to default configuration.")
		return defaultConfig
	}

	// If there are any nil values in uncheckedConfig, use defaultConfig
	invalid := uncheckedConfig.NumRuns == nil || uncheckedConfig.SelectionMethod == nil || uncheckedConfig.PopulationSize == nil || uncheckedConfig.MaxGenerations == nil || uncheckedConfig.NumQueens == nil || uncheckedConfig.MutationRate == nil || uncheckedConfig.CrossOverRate == nil || uncheckedConfig.Elitism == nil
	if invalid {
		fmt.Println("[WARN] Invalid config file, falling back to default configuration.")
		return defaultConfig
	}

	config := Config{
		NumRuns:         *uncheckedConfig.NumRuns,
		SelectionMethod: *uncheckedConfig.SelectionMethod,
		PopulationSize:  *uncheckedConfig.PopulationSize,
		MaxGenerations:  *uncheckedConfig.MaxGenerations,
		NumQueens:       *uncheckedConfig.NumQueens,
		MutationRate:    *uncheckedConfig.MutationRate,
		CrossOverRate:   *uncheckedConfig.CrossOverRate,
		Elitism:         *uncheckedConfig.Elitism,
	}

	return config
}
