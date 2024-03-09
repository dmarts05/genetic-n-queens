package main

import (
	"flag"
	"fmt"
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
	config := config.LoadConfig(filepath.Join(path))

	fmt.Println("************************************************************")
	fmt.Println("Starting genetic algorithm with the following configuration:")
	fmt.Println("************************************************************")
	fmt.Println("\t- Number of runs: ", config.NumRuns)
	fmt.Println("\t- Selection method: ", config.SelectionMethod)
	fmt.Println("\t- Population size: ", config.PopulationSize)
	fmt.Println("\t- Maximum number of generations: ", config.MaxGenerations)
	fmt.Println("\t- Number of queens: ", config.NumQueens)
	fmt.Println("\t- Mutation rate: ", config.MutationRate)
	fmt.Println("\t- Crossover rate: ", config.CrossOverRate)
	fmt.Println("\t- Elitism: ", config.Elitism)
	fmt.Println("************************************************************")
}
