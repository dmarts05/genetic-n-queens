from deap import tools
from deap.algorithms import varAnd


def eaSimpleFitnessStop(
    population,
    toolbox,
    cxpb: float,
    mutpb: float,
    ngen: int,
    best_possible_fitness: int,
    elitism: bool = False,
    stats=None,
    halloffame=None,
    verbose=__debug__,
):
    logbook = tools.Logbook()
    logbook.header = ["gen", "nevals"] + (stats.fields if stats else [])  # type: ignore

    # Evaluate the individuals with an invalid fitness
    invalid_ind = [ind for ind in population if not ind.fitness.valid]
    fitnesses = toolbox.map(toolbox.evaluate, invalid_ind)
    for ind, fit in zip(invalid_ind, fitnesses):
        ind.fitness.values = fit

    if halloffame is not None:
        halloffame.update(population)

    record = stats.compile(population) if stats else {}
    logbook.record(gen=0, nevals=len(invalid_ind), **record)
    if verbose:
        print(logbook.stream)

    # Begin the generational process
    for gen in range(1, ngen + 1):
        # Stop if the best individual has the best possible fitness
        best_current_fitness = int(halloffame[0].fitness.values[0]) if halloffame else 0
        if best_current_fitness == best_possible_fitness:
            break

        # Select the next generation individuals
        offspring = toolbox.select(population, len(population))

        # Vary the pool of individuals
        offspring = varAnd(offspring, toolbox, cxpb, mutpb)

        # Evaluate the individuals with an invalid fitness
        invalid_ind = [ind for ind in offspring if not ind.fitness.valid]
        fitnesses = toolbox.map(toolbox.evaluate, invalid_ind)
        for ind, fit in zip(invalid_ind, fitnesses):
            ind.fitness.values = fit

        # Update the hall of fame with the generated individuals
        if halloffame is not None:
            halloffame.update(offspring)

        # Replace the current population by the offspring
        population[:] = offspring

        # Perform elitist reduction if enabled
        if elitism:
            extendedPopulation = population + offspring
            elites = tools.selBest(extendedPopulation, len(population))
            population = elites

        # Append the current generation statistics to the logbook
        record = stats.compile(population) if stats else {}
        logbook.record(gen=gen, nevals=len(invalid_ind), **record)
        if verbose:
            print(logbook.stream)

        gen += 1

    return population, logbook
