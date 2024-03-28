package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/dmarts05/genetic-n-queens/internal/config"
	"github.com/dmarts05/genetic-n-queens/internal/population"
	"github.com/dmarts05/genetic-n-queens/internal/result"
)

func main() {
	// Load config
	var help bool
	var configPath string

	flag.BoolVar(&help, "h", false, "Show help")
	flag.StringVar(&configPath, "c", "", "Provide the path to a JSON configuration file for the genetic algorithm.")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	bestPossibleFitness := cfg.NumQueens * (cfg.NumQueens - 1) / 2

	fmt.Println("************************************************************")
	fmt.Println("Starting genetic algorithm with the following configuration:")
	fmt.Println("- Number of runs:", cfg.NumRuns)
	fmt.Println("- Selection method:", cfg.SelectionMethod)
	fmt.Println("- Population size:", cfg.PopulationSize)
	fmt.Println("- Maximum number of generations:", cfg.MaxGenerations)
	fmt.Println("- Number of queens:", cfg.NumQueens)
	fmt.Println("- Mutation rate:", cfg.MutationRate)
	fmt.Println("- Crossover rate:", cfg.CrossOverRate)
	fmt.Println("- Elitism:", cfg.Elitism)
	fmt.Println("- Best possible fitness:", bestPossibleFitness)
	fmt.Println("************************************************************")

	fmt.Println()

	// Run the genetic algorithm for the number of runs specified in the configuration with goroutines
	var wg sync.WaitGroup
	ch := make(chan result.GenerationResult, cfg.NumRuns)
	for i := 0; i < cfg.NumRuns; i++ {
		pop := population.Generate(cfg.NumQueens, cfg.PopulationSize)
		wg.Add(1)
		go population.EvolveConcurrentWrapper(i+1, ch, &wg, pop, cfg.SelectionMethod, cfg.MaxGenerations, cfg.MutationRate, cfg.CrossOverRate, cfg.Elitism, bestPossibleFitness)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(ch)

	// Load results from the channel
	results := []result.GenerationResult{}
	for r := range ch {
		results = append(results, r)
	}

	fmt.Println()

	// Show final results
	fmt.Println("************************************************************")
	fmt.Println("Final results:")
	fmt.Println("- Mean number of generations:", result.GetMeanGenerations(results))
	fmt.Println("- Best fitness:", result.GetBestFitness(results))
	fmt.Println("- Worst fitness:", result.GetWorstFitness(results))
	fmt.Println("- Mean of the best fitness:", result.GetMeanBestFitness(results))
	fmt.Println("- Mean of the mean fitness:", result.GetMeanMeanFitness(results))
	fmt.Println("************************************************************")

	// Save results to a file
	fileName := "results.json"
	err = result.SaveResultsToFile(results, fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Results saved to:", fileName)
}
