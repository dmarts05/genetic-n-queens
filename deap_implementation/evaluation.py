def get_num_clashes(individual: list[int]) -> int:
    """Get the number of queen clashes in the individual."""

    # We only need to check for diagonal clashes because we are using permutations
    num_queens = len(individual)
    clashes = 0
    for col1 in range(num_queens):
        for col2 in range(col1 + 1, num_queens):
            row1 = individual[col1]
            row2 = individual[col2]
            if row1 - col1 == row2 - col2 or row1 + col1 == row2 + col2:
                clashes += 1
    return clashes


def eval_fitness(individual: list[int]) -> tuple[int,]:
    """Evaluate the fitness of the individual."""
    num_queens = len(individual)
    max_non_attacking_pairs = int(num_queens * (num_queens - 1) / 2)
    clashes = get_num_clashes(individual)
    fitness = max_non_attacking_pairs - clashes
    return (fitness,)
