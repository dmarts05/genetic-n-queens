package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
)

var DefaultConfig = Config{
	NumRuns:         12,
	SelectionMethod: Tournament,
	PopulationSize:  300,
	MaxGenerations:  3000,
	NumQueens:       29,
	MutationRate:    0.2,
	CrossOverRate:   0.5,
	Elitism:         false,
	TournamentSize:  3,
}

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
	TournamentSize  int                 `json:"tournament_size"`
}

func New(selectionMethod SelectionMethodType, tournamentSize, numRuns, populationSize, maxGenerations, numQueens int, mutationRate, crossOverRate float64, elitism bool) (Config, error) {
	if selectionMethod != Tournament {
		tournamentSize = 0
	}

	cfg := Config{
		SelectionMethod: selectionMethod,
		NumRuns:         numRuns,
		PopulationSize:  populationSize,
		MaxGenerations:  maxGenerations,
		NumQueens:       numQueens,
		MutationRate:    mutationRate,
		CrossOverRate:   crossOverRate,
		Elitism:         elitism,
		TournamentSize:  tournamentSize,
	}
	err := cfg.validate()
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// Validate configuration values
func (c Config) validate() error {
	switch {
	case c.NumRuns < 1:
		return errors.New("number of runs must be at least 1")
	case c.PopulationSize < 1 || c.PopulationSize%2 != 0:
		return errors.New("population size must be at least 2 and even")
	case c.MaxGenerations < 1:
		return errors.New("maximum number of generations must be at least 1")
	case c.NumQueens < 4:
		return errors.New("number of queens must be at least 4, otherwise the problem is trivial")
	case c.MutationRate < 0 || c.MutationRate > 1:
		return errors.New("mutation rate must be between 0 and 1")
	case c.CrossOverRate < 0 || c.CrossOverRate > 1:
		return errors.New("crossover rate must be between 0 and 1")
	case c.SelectionMethod == Tournament && c.TournamentSize < 2:
		return errors.New("tournament size must be at least 2 when using the tournament selection method")
	default:
		return nil
	}
}

// Load configuration from specified path or use default configuration if not found
func LoadConfigFromJSON(path string) (Config, error) {
	// Load config from specified path
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("load config: specified configuration file not found, falling back to default configuration")
		return DefaultConfig, nil
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
		TournamentSize  int                  `json:"tournament_size"`
	}{}

	// Load json into uncheckedConfig
	err = json.Unmarshal(data, &uncheckedConfig)
	if err != nil {
		return Config{}, fmt.Errorf("load config: invalid config file: %w", err)
	}

	// Loop through uncheckedConfig and check if any of the fields are nil
	v := reflect.ValueOf(uncheckedConfig)
	for i := 0; i < v.NumField(); i++ {
		// TournamentSize is not a pointer, so we need to check it separately
		if i == 8 {
			continue
		}

		if v.Field(i).IsNil() {
			return Config{}, errors.New("load config: invalid config file")
		}
	}

	// Since there are no nil fields, we can safely load the values into the config struct
	var cfg Config
	_ = json.Unmarshal(data, &cfg)
	err = cfg.validate()
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
