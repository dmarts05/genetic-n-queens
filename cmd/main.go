package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/dmarts05/genetic-n-queens/internal/config"
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
}
