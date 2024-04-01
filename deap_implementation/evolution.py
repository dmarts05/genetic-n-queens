import argparse
from concurrent.futures import ThreadPoolExecutor
import random

from deap import base, creator, tools
import numpy
from deap_implementation import custom_algorithms
from deap_implementation.evaluation import eval_fitness
from deap_implementation.schemas import Result, SelectionMethod


def set_up_types() -> None:
    """Set up DEAP types for the N-Queens problem."""
    creator.create("FitnessMax", base.Fitness, weights=(1.0,))
    creator.create("Individual", list, fitness=creator.FitnessMax)  # type: ignore


def set_up_toolbox(args: argparse.Namespace) -> base.Toolbox:
    """Set up DEAP toolbox for the N-Queens problem."""
    toolbox = base.Toolbox()

    # Individuals are lists of integers from 0 to num_queens - 1
    # Each index represents a column and the value represents the row
    # This makes permutations a good representation for the problem because we can guarantee that each row and column will have exactly one queen
    toolbox.register(
        "permutation", random.sample, range(args.num_queens), args.num_queens
    )
    toolbox.register(
        "individual",
        tools.initIterate,
        creator.Individual,  # type: ignore
        toolbox.permutation,  # type: ignore
    )
    toolbox.register("population", tools.initRepeat, list, toolbox.individual)  # type: ignore

    toolbox.register("evaluate", eval_fitness)

    # We use ordered crossover because it guarantees that each row and column will have exactly one queen. PMX could also work but it is more complex
    toolbox.register("mate", tools.cxOrdered)

    # As we increase the number of queens, the probability of mutating each individual should decrease because the search space gets broader
    individual_mutate_prob = 2.0 / args.num_queens
    toolbox.register("mutate", tools.mutShuffleIndexes, indpb=individual_mutate_prob)

    if args.selection_method == SelectionMethod.ROULETTE.value:
        toolbox.register("select", tools.selRoulette)
    elif args.selection_method == SelectionMethod.TOURNAMENT.value:
        toolbox.register("select", tools.selTournament, tournsize=args.tournament_size)

    return toolbox


def evolve_concurrent_wrapper(
    toolbox,
    num_runs: int,
    population_size: int,
    max_generations: int,
    mutation_rate: float,
    crossover_rate: float,
    best_possible_fitness: int,
    elitism: bool = False,
    verbose: bool = False,
) -> list[Result]:
    results: list[Result] = []

    hof = tools.HallOfFame(1)
    stats = tools.Statistics(lambda ind: ind.fitness.values)
    stats.register("Avg", numpy.mean)
    stats.register("Min", numpy.min)
    stats.register("Max", numpy.max)

    with ThreadPoolExecutor() as executor:
        futures = [
            executor.submit(
                evolve,
                toolbox,
                stats,
                hof,
                population_size,
                max_generations,
                mutation_rate,
                crossover_rate,
                best_possible_fitness,
                elitism,
                verbose,
            )
            for _ in range(num_runs)
        ]

        for i, future in enumerate(futures):
            result = future.result()
            results.append(result)
            print("-" * 60)
            if result.best_fitness == best_possible_fitness:
                print(
                    f"Worker {i + 1} has found one of the optimal solutions: {result.best_queen_positions}"
                )
            else:
                print(
                    f"Worker {i + 1} has finished with a suboptimal solution: {result.best_queen_positions} with fitness {result.best_fitness}"
                )
            print("-" * 60)
            print()

    return results


def evolve(
    toolbox,
    stats,
    hof,
    population_size: int,
    max_generations: int,
    mutation_rate: float,
    crossover_rate: float,
    best_possible_fitness: int,
    elitism: bool = False,
    verbose: bool = False,
) -> Result:
    pop = toolbox.population(n=population_size)
    pop, logbook = custom_algorithms.eaSimpleFitnessStop(
        pop,
        toolbox,
        cxpb=crossover_rate,
        mutpb=mutation_rate,
        ngen=max_generations,
        stats=stats,
        halloffame=hof,
        verbose=verbose,
        best_possible_fitness=best_possible_fitness,
        elitism=elitism,
    )

    best_fitness = int(hof[0].fitness.values[0])
    result = Result(
        best_fitness=best_fitness,
        best_queen_positions=hof[0],
        generation=logbook.select("gen")[-1],  # type: ignore
        is_solution=best_fitness == best_possible_fitness,
        mean_fitness=logbook.select("Avg")[-1],  # type: ignore
    )

    return result
