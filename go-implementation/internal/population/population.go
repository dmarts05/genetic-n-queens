package population

import (
	"fmt"
	"log"
	"math/rand/v2"
	"sync"

	"github.com/dmarts05/genetic-n-queens/internal/config"
	"github.com/dmarts05/genetic-n-queens/internal/individual"
	"github.com/dmarts05/genetic-n-queens/internal/result"
	"github.com/dmarts05/genetic-n-queens/internal/selection"
)

const tournamentSize = 3

// Generate a random individual
func generateRandomIndividual(numQueens int) *individual.Individual {
	return &individual.Individual{QueenPositions: rand.Perm(numQueens)}
}

// Generate a population of random individuals with the given number of queens and population size
func Generate(numQueens, populationSize int) []*individual.Individual {
	population := make([]*individual.Individual, populationSize)

	for i := 0; i < populationSize; i++ {
		population[i] = generateRandomIndividual(numQueens)
	}

	return population
}

// Wrapper for Evolve function to be used with goroutines
func EvolveConcurrentWrapper(workerID int, ch chan<- result.GenerationResult, wg *sync.WaitGroup, pop []*individual.Individual, selectionMethod config.SelectionMethodType, maxGenerations int, mutationRate float64, crossoverRate float64, elitism bool, bestPossibleFitness int) {
	var r result.GenerationResult

	defer func() {
		fmt.Println("------------------------------------------------------------")
		if r.BestFitness == bestPossibleFitness {
			fmt.Println("Worker", workerID, "has found one of the optimal solutions:", r.BestQueenPositions)
		} else {
			fmt.Println("Worker", workerID, "has finished with a suboptimal solution:", r.BestQueenPositions, "with fitness", r.BestFitness)
		}
		fmt.Println("------------------------------------------------------------")
		wg.Done()
	}()

	r = Evolve(pop, selectionMethod, maxGenerations, mutationRate, crossoverRate, elitism, bestPossibleFitness)
	ch <- r
}

// Evolve the population by applying the selection, crossover and mutation methods and return the best generation
func Evolve(pop []*individual.Individual, selectionMethod config.SelectionMethodType, maxGenerations int, mutationRate float64, crossoverRate float64, elitism bool, bestPossibleFitness int) result.GenerationResult {
	results := []result.GenerationResult{}

	for generation := 1; generation <= maxGenerations; generation++ {
		// Evaluate fitness
		bestIndividual := pop[0]
		for _, ind := range pop {
			fitness := ind.Fitness()
			if fitness > bestIndividual.Fitness() {
				bestIndividual = ind
			}
		}

		meanFitness := 0.0
		for _, ind := range pop {
			meanFitness += float64(ind.Fitness())
		}
		meanFitness = meanFitness / float64(len(pop))

		bestQueenPositions := make([]int, len(bestIndividual.QueenPositions))
		copy(bestQueenPositions, bestIndividual.QueenPositions)
		bestFitness := bestIndividual.Fitness()
		results = append(results, result.GenerationResult{
			Generation:         generation,
			BestQueenPositions: bestQueenPositions,
			BestFitness:        bestFitness,
			MeanFitness:        meanFitness,
			IsSolution:         bestFitness == bestPossibleFitness,
		})

		// Check if we have reached the best possible fitness
		if bestIndividual.Fitness() == bestPossibleFitness {
			break
		}

		// Select parents
		var parents []*individual.Individual
		switch selectionMethod {
		case config.Roulette:
			parents = selection.SelectByRoulette(pop)
		case config.Tournament:
			parents = selection.SelectByTournament(pop, tournamentSize)
		}

		// Crossover
		newPop := []*individual.Individual{}
		for i := 0; i < len(parents); i += 2 {
			doCrossover := rand.Float64() < crossoverRate
			if i+1 < len(parents) {
				parent1 := parents[i]
				parent2 := parents[i+1]
				if doCrossover {
					child1, child2, err := parent1.Crossover(parent2)
					if err != nil {
						log.Fatal(err)
					}
					newPop = append(newPop, child1, child2)
				} else {
					newPop = append(newPop, parent1, parent2)
				}
			}
		}

		// Mutate
		for _, ind := range newPop {
			doMutate := rand.Float64() < mutationRate
			if doMutate {
				// Since the mutation rate is per individual, we need to adjust it based on the number of queens
				numQueens := len(ind.QueenPositions)
				ind.Mutate(2.0 / float64(numQueens))
			}
		}

		// Perform elitist reduction if enabled
		if elitism {
			extendedPop := append(pop, newPop...)
			elites := selection.SelectByElitism(extendedPop, len(pop))
			newPop = elites
		}

		pop = newPop
	}

	// Get best result
	best_result := results[0]
	for _, result := range results {
		if result.BestFitness > best_result.BestFitness {
			best_result = result
		}
	}

	return best_result
}
