import argparse
import json
import time
from deap_implementation.schemas import SelectionMethod
from deap_implementation.evolution import (
    evolve_concurrent_wrapper,
    set_up_toolbox,
    set_up_types,
)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="N-Queens DEAP implementation")
    parser.add_argument(
        "-numRuns", type=int, default=12, help="Number of runs", dest="num_runs"
    )
    parser.add_argument(
        "-populationSize",
        type=int,
        default=300,
        help="Population size",
        dest="population_size",
    )
    parser.add_argument(
        "-maxGenerations",
        type=int,
        default=3000,
        help="Maximum number of generations",
        dest="max_generations",
    )
    parser.add_argument(
        "-numQueens", type=int, default=29, help="Number of queens", dest="num_queens"
    )
    parser.add_argument(
        "-mutationRate",
        type=float,
        default=0.2,
        help="Mutation rate",
        dest="mutation_rate",
    )
    parser.add_argument(
        "-crossOverRate",
        type=float,
        default=0.5,
        help="Crossover rate",
        dest="crossover_rate",
    )
    parser.add_argument(
        "-selectionMethod",
        choices=[SelectionMethod.TOURNAMENT.value, SelectionMethod.ROULETTE.value],
        help="Selection method",
        default=SelectionMethod.TOURNAMENT.value,
        dest="selection_method",
    )
    parser.add_argument(
        "-tournamentSize",
        type=int,
        default=3,
        help="Tournament size",
        dest="tournament_size",
    )
    parser.add_argument(
        "-elitism",
        type=bool,
        default=False,
        action=argparse.BooleanOptionalAction,
        help="Elitism",
        dest="elitism",
    )

    args = parser.parse_args()
    return args


def validate_args(args: argparse.Namespace) -> None:
    if (
        args.selection_method == SelectionMethod.TOURNAMENT.value
        and args.tournament_size < 2
    ):
        raise ValueError("Tournament size must be at least 2")

    if (
        args.selection_method == SelectionMethod.ROULETTE.value
        and args.population_size % 2 != 0
    ):
        raise ValueError("Population size must be even when using roulette selection")

    if args.mutation_rate < 0 or args.mutation_rate > 1:
        raise ValueError("Mutation rate must be between 0 and 1")

    if args.crossover_rate < 0 or args.crossover_rate > 1:
        raise ValueError("Crossover rate must be between 0 and 1")

    if args.num_queens < 4:
        raise ValueError("Number of queens must be at least 4")

    if args.max_generations < 1:
        raise ValueError("Maximum number of generations must be at least 1")


def main() -> None:
    """Run the DEAP implementation of the N-Queens problem."""
    args = parse_args()
    validate_args(args)

    best_possible_fitness = int(args.num_queens * (args.num_queens - 1) / 2)
    print("*" * 60)
    print("Starting genetic algorithm with the following configuration:")
    print(f"- Number of runs: {args.num_runs}")
    print(f"- Selection method: {args.selection_method}")
    print(f"- Population size: {args.population_size}")
    print(f"- Maximum number of generations: {args.max_generations}")
    print(f"- Number of queens: {args.num_queens}")
    print(f"- Mutation rate: {args.mutation_rate}")
    print(f"- Crossover rate: {args.crossover_rate}")
    print(f"- Elitism: {args.elitism}")
    print(f"- Best possible fitness: {best_possible_fitness}")
    print("*" * 60)

    print()

    set_up_types()
    toolbox = set_up_toolbox(args)

    start = time.time()

    results = evolve_concurrent_wrapper(
        toolbox=toolbox,
        num_runs=args.num_runs,
        population_size=args.population_size,
        max_generations=args.max_generations,
        mutation_rate=args.mutation_rate,
        crossover_rate=args.crossover_rate,
        best_possible_fitness=best_possible_fitness,
        elitism=args.elitism,
    )

    elapsed = time.time() - start

    print("*" * 60)
    print("Final results:")
    print(f"- Elapsed time: {elapsed} seconds")
    print(
        f"- Number of solutions found: {len([result for result in results if result.is_solution])}"
    )
    print(
        f"- Mean number of generations: {sum([result.generation for result in results]) / len(results)}"
    )
    print(f"- Best fitness: {max([result.best_fitness for result in results])}")
    print(f"- Worst fitness: {min([result.best_fitness for result in results])}")
    print(
        f"- Mean of the best fitness: {sum([result.best_fitness for result in results]) / len(results)}"
    )
    print(
        f"- Mean of the mean fitness: {sum([result.mean_fitness for result in results]) / len(results)}"
    )
    print("*" * 60)

    # Save results to a file
    with open("results.json", "w") as f:
        json.dump(
            [
                {
                    "best_queen_positions": result.best_queen_positions,
                    "generation": result.generation,
                    "best_fitness": result.best_fitness,
                    "mean_fitness": result.mean_fitness,
                    "is_solution": result.is_solution,
                }
                for result in results
            ],
            f,
            indent=4,
        )


if __name__ == "__main__":
    main()
