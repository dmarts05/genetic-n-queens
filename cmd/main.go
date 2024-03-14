package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/dmarts05/genetic-n-queens/internal/config"
	"github.com/dmarts05/genetic-n-queens/internal/population"
	"github.com/dmarts05/genetic-n-queens/internal/result"
)

func main() {
	// Load config
	var help bool
	var configName string

	flag.BoolVar(&help, "h", false, "Show help")
	flag.StringVar(&configName, "c", "", "Provide the name of the configuration file you want to use in configs folder.")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	path := filepath.Join("configs", configName)
	config, err := config.LoadConfig(filepath.Join(path))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("************************************************************")
	fmt.Println("Starting genetic algorithm with the following configuration:")
	fmt.Println("************************************************************")
	fmt.Println("- Number of runs: ", config.NumRuns)
	fmt.Println("- Selection method: ", config.SelectionMethod)
	fmt.Println("- Population size: ", config.PopulationSize)
	fmt.Println("- Maximum number of generations: ", config.MaxGenerations)
	fmt.Println("- Number of queens: ", config.NumQueens)
	fmt.Println("- Mutation rate: ", config.MutationRate)
	fmt.Println("- Crossover rate: ", config.CrossOverRate)
	fmt.Println("- Elitism: ", config.Elitism)
	fmt.Println("************************************************************")

	bestPossibleFitness := config.NumQueens * (config.NumQueens - 1) / 2
	results := []result.GenerationResult{}
	pop := population.GeneratePopulation(config.NumQueens, config.PopulationSize)

	for generation := 1; generation <= config.MaxGenerations; generation++ {
		fmt.Println("----------------------------------------------------------")
		fmt.Println("Generation: ", generation)
		fmt.Println("----------------------------------------------------------")

		// Evaluate fitness
		best_individual := pop[0]
		for _, ind := range pop {
			fitness := ind.Fitness()
			if fitness > best_individual.Fitness() {
				best_individual = ind
			}
		}

		mean_fitness := 0.0
		for _, ind := range pop {
			mean_fitness += float64(ind.Fitness())
		}
		mean_fitness = mean_fitness / float64(len(pop))

		fmt.Println("Best fitness: ", best_individual.Fitness())
		fmt.Println("Mean fitness: ", mean_fitness)

		results = append(results, result.GenerationResult{
			Generation:         generation,
			BestQueenPositions: best_individual.QueenPositions,
			BestFitness:        best_individual.Fitness(),
			MeanFitness:        mean_fitness,
		})

		// Check if we have reached the best possible fitness
		if best_individual.Fitness() == bestPossibleFitness {
			break
		}
	}

	fmt.Println(results)
}
