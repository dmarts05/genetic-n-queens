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
	flag.StringVar(&configPath, "c", "", "Provide the name of the configuration file you want to use in configs folder.")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("************************************************************")
	fmt.Println("Starting genetic algorithm with the following configuration:")
	fmt.Println("- Number of runs: ", cfg.NumRuns)
	fmt.Println("- Selection method: ", cfg.SelectionMethod)
	fmt.Println("- Population size: ", cfg.PopulationSize)
	fmt.Println("- Maximum number of generations: ", cfg.MaxGenerations)
	fmt.Println("- Number of queens: ", cfg.NumQueens)
	fmt.Println("- Mutation rate: ", cfg.MutationRate)
	fmt.Println("- Crossover rate: ", cfg.CrossOverRate)
	fmt.Println("- Elitism: ", cfg.Elitism)
	fmt.Println("************************************************************")

	// Run the genetic algorithm for the number of runs specified in the configuration with goroutines
	bestPossibleFitness := cfg.NumQueens * (cfg.NumQueens - 1) / 2
	var wg sync.WaitGroup
	ch := make(chan result.GenerationResult, cfg.NumRuns)
	for i := 0; i < cfg.NumRuns; i++ {
		pop := population.Generate(cfg.NumQueens, cfg.PopulationSize)
		wg.Add(1)
		go population.EvolveConcurrentWrapper(i, ch, &wg, pop, cfg.SelectionMethod, cfg.MaxGenerations, cfg.MutationRate, cfg.CrossOverRate, cfg.Elitism, bestPossibleFitness)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(ch)

	// Load results from the channel
	results := []result.GenerationResult{}
	for r := range ch {
		results = append(results, r)
	}

	// Save results to a file
	err = result.SaveResultsToFile(results, "results.json")
	if err != nil {
		log.Fatal(err)
	}
}
