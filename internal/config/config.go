package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
)

// Represents the available selection methods for the genetic algorithm
type SelectionMethodType string

const (
	Tournament SelectionMethodType = "tournament"
	Roulette   SelectionMethodType = "roulette"
)

// Represents the configuration for the genetic algorithm
type Config struct {
	SelectionMethod SelectionMethodType `json:"selection_method"`
	NumRuns         int                 `json:"num_runs"`
	PopulationSize  int                 `json:"population_size"`
	MaxGenerations  int                 `json:"max_generations"`
	NumQueens       int                 `json:"num_queens"`
	MutationRate    float64             `json:"mutation_rate"`
	CrossOverRate   float64             `json:"crossover_rate"`
	Elitism         bool                `json:"elitism"`
}

// Load configuration from specified path or use default configuration if not found
func LoadConfig(path string) (Config, error) {
	defaultConfig := Config{
		NumRuns:         5,
		SelectionMethod: Tournament,
		PopulationSize:  16,
		MaxGenerations:  500,
		NumQueens:       8,
		MutationRate:    0.2,
		CrossOverRate:   0.5,
		Elitism:         false,
	}

	// Load config from specified path
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("load config: specified configuration file not found, falling back to default configuration")
		return defaultConfig, nil
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
		return Config{}, fmt.Errorf("load config: invalid config file: %w", err)
	}

	// Loop through uncheckedConfig and check if any of the fields are nil
	v := reflect.ValueOf(uncheckedConfig)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsNil() {
			return Config{}, errors.New("load config: invalid config file")
		}
	}

	// Check config values are valid
	switch {
	case *uncheckedConfig.NumRuns < 1:
		return Config{}, errors.New("load config: number of runs must be at least 1")
	case *uncheckedConfig.PopulationSize < 1 || *uncheckedConfig.PopulationSize%2 != 0:
		return Config{}, errors.New("load config: population size must be at least 2 and even")
	case *uncheckedConfig.MaxGenerations < 1:
		return Config{}, errors.New("load config: maximum number of generations must be at least 1")
	case *uncheckedConfig.NumQueens < 4:
		return Config{}, errors.New("load config: number of queens must be at least 4, otherwise the problem is trivial")
	case *uncheckedConfig.MutationRate < 0 || *uncheckedConfig.MutationRate > 1:
		return Config{}, errors.New("load config: mutation rate must be between 0 and 1")
	case *uncheckedConfig.CrossOverRate < 0 || *uncheckedConfig.CrossOverRate > 1:
		return Config{}, errors.New("load config: crossover rate must be between 0 and 1")
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

	return config, nil
}
