from dataclasses import dataclass
from enum import Enum


@dataclass
class Result:
    """Stores the result of a run of the genetic algorithm."""

    generation: int
    best_queen_positions: list[int]
    best_fitness: int
    mean_fitness: float
    is_solution: bool


class SelectionMethod(Enum):
    """Represents the selection method to use in the genetic algorithm."""

    TOURNAMENT = "tournament"
    ROULETTE = "roulette"
