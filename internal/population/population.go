package population

import (
	"fmt"
	"log"
	"math/rand/v2"
	"sync"

	"github.com/dmarts05/genetic-n-queens/internal/config"
	"github.com/dmarts05/genetic-n-queens/internal/individual"
	"github.com/dmarts05/genetic-n-queens/internal/position"
	"github.com/dmarts05/genetic-n-queens/internal/result"
	"github.com/dmarts05/genetic-n-queens/internal/selection"
)

// Generate a random individual
func generateRandomIndividual(numQueens int) *individual.Individual {
	queenPositions := make([]position.Position, numQueens)
	for i := 0; i < numQueens; i++ {
		// We only make the column random in order to avoid having queens in the same row, reducing the number of clashes by default
		queenPositions[i] = position.Position{
			Row:    i,
			Column: rand.IntN(numQueens),
		}
	}
	return individual.New(queenPositions)
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
		fmt.Println("Worker", workerID, "finished with best fitness: ", r.BestFitness)
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

		bestQueenPositions := make([]position.Position, len(bestIndividual.QueenPositions))
		copy(bestQueenPositions, bestIndividual.QueenPositions)
		results = append(results, result.GenerationResult{
			Generation:         generation,
			BestQueenPositions: bestQueenPositions,
			BestFitness:        bestIndividual.Fitness(),
			MeanFitness:        meanFitness,
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
			parents = selection.SelectByTournament(pop)
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
				ind.Mutate()
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
