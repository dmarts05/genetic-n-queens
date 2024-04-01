import argparse
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
        help="Elitism",
        dest="elitism",
    )

    args = parser.parse_args()
    return args


def main() -> None:
    """Run the DEAP implementation of the N-Queens problem."""
    args = parse_args()

    set_up_types()
    toolbox = set_up_toolbox(args)

    best_possible_fitness = int(args.num_queens * (args.num_queens - 1) / 2)
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
    print(results)


if __name__ == "__main__":
    main()
