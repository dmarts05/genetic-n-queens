import argparse
import json
import tkinter as tk

CANVAS_WIDTH = 400
CANVAS_HEIGHT = 400
CELL_SIZE = 50
QUEEN_SIZE = 20


class Solution:
    def __init__(
        self,
        best_queen_positions: list[int],
        generation: int,
        best_fitness: int,
        mean_fitness: float,
    ):
        self.best_queen_positions = best_queen_positions
        self.generation = generation
        self.best_fitness = best_fitness
        self.mean_fitness = mean_fitness


class SolutionCarousel(tk.Tk):
    def __init__(
        self,
        width: int,
        height: int,
        cell_size: int,
        queen_size: int,
        results: list[Solution],
    ) -> None:
        if not results:
            raise ValueError("Results list cannot be empty")

        super().__init__()

        self.width = width
        self.height = height
        self.cell_size = cell_size
        self.queen_size = queen_size
        self.num_queens = len(results[0].best_queen_positions)

        self.title("N-Queens Solution Carousel")

        # Add padding
        self.padding = 10

        # Create frames for better organization
        self.canvas_frame = tk.Frame(self)
        self.canvas_frame.pack(
            fill=tk.BOTH, expand=True, padx=self.padding, pady=(self.padding, 0)
        )

        self.solution_frame = tk.Frame(self)
        self.solution_frame.pack(fill=tk.X, padx=self.padding, pady=(self.padding, 0))

        self.button_frame = tk.Frame(self)
        self.button_frame.pack(fill=tk.X, padx=self.padding, pady=(0, self.padding))

        # Canvas for drawing chessboard and queens
        self.canvas = tk.Canvas(self.canvas_frame, width=self.width, height=self.height)
        self.canvas.pack(padx=(0, 0), pady=(0, 0))

        # Solution info label
        self.solution_label = tk.Label(self.solution_frame, font=("Arial", 12))
        self.solution_label.pack(fill=tk.X)

        self.results = results
        self.current_index = 0

        # Create buttons
        self.prev_button = tk.Button(
            self.button_frame, text="<<", command=self.show_previous
        )
        self.prev_button.pack(side="left", padx=(0, 5))

        self.next_button = tk.Button(
            self.button_frame, text=">>", command=self.show_next
        )
        self.next_button.pack(side="right", padx=(5, 0))

        self.create_widgets()

    def create_widgets(self) -> None:
        self.draw_chessboard()
        self.draw_queens(self.current_index)
        self.update_solution_info()
        self.update_button_states()

    def draw_chessboard(self) -> None:
        for i in range(self.num_queens):
            for j in range(self.num_queens):
                x0, y0 = j * self.cell_size, i * self.cell_size
                x1, y1 = x0 + self.cell_size, y0 + self.cell_size
                color = "white" if (i + j) % 2 == 0 else "light grey"
                self.canvas.create_rectangle(x0, y0, x1, y1, fill=color)

    def draw_queens(self, index: int) -> None:
        queens_positions = self.results[index].best_queen_positions
        for i, row in enumerate(queens_positions):
            x0, y0 = row * self.cell_size, i * self.cell_size
            self.canvas.create_text(
                x0 + self.cell_size // 2,
                y0 + self.cell_size // 2,
                text="Q",
                font=("Arial", self.queen_size),
            )

        # Highlight conflicts
        for i, row in enumerate(queens_positions):
            for j, other_row in enumerate(queens_positions):
                if i == j:
                    continue
                if row == other_row or abs(row - other_row) == abs(i - j):
                    x0, y0 = row * self.cell_size, i * self.cell_size
                    x1, y1 = other_row * self.cell_size, j * self.cell_size
                    self.canvas.create_line(
                        x0 + self.cell_size // 2,
                        y0 + self.cell_size // 2,
                        x1 + self.cell_size // 2,
                        y1 + self.cell_size // 2,
                        fill="red",
                        width=2,
                    )

    def update_solution_info(self) -> None:
        current_solution = self.results[self.current_index]
        solution_info = f"Solution {self.current_index + 1}/{len(self.results)}\n"
        solution_info += "-" * 40 + "\n"
        solution_info += f"Queens: {self.num_queens}\n"
        solution_info += f"Generation: {current_solution.generation}\n"
        solution_info += f"Best Fitness: {current_solution.best_fitness}\n"
        solution_info += f"Mean Fitness: {current_solution.mean_fitness}\n"
        solution_info += f"Conflicts: {self.calculate_num_conflicts(current_solution.best_queen_positions)}\n"
        self.solution_label.config(text=solution_info)

    def update_button_states(self) -> None:
        if self.current_index == 0:
            self.prev_button.config(state="disabled")
        else:
            self.prev_button.config(state="normal")

        if self.current_index == len(self.results) - 1:
            self.next_button.config(state="disabled")
        else:
            self.next_button.config(state="normal")

    def show_previous(self) -> None:
        self.current_index -= 1
        if self.current_index < 0:
            self.current_index = 0
        self.redraw()
        self.update_button_states()

    def show_next(self) -> None:
        self.current_index += 1
        if self.current_index >= len(self.results):
            self.current_index = len(self.results) - 1
        self.redraw()
        self.update_button_states()

    def redraw(self) -> None:
        self.canvas.delete("all")
        self.draw_chessboard()
        self.draw_queens(self.current_index)
        self.update_solution_info()

    def calculate_num_conflicts(self, queen_positions: list[int]) -> int:
        # Only check pairs of queens once
        num_conflicts = 0
        for i in range(len(queen_positions)):
            for j in range(i + 1, len(queen_positions)):
                if queen_positions[i] == queen_positions[j] or abs(
                    queen_positions[i] - queen_positions[j]
                ) == abs(i - j):
                    num_conflicts += 1

        return num_conflicts


def load_results_from_json(json_path: str) -> list[Solution]:
    with open(json_path, "r") as f:
        data = json.load(f)
    return [Solution(**item) for item in data]


def main() -> None:
    parser = argparse.ArgumentParser(description="Showcase N-Queens results.")
    parser.add_argument(
        "json_path",
        type=str,
        help="Path to the JSON file containing the results.",
    )
    args = parser.parse_args()

    results = load_results_from_json(args.json_path)
    app = SolutionCarousel(CANVAS_WIDTH, CANVAS_HEIGHT, CELL_SIZE, QUEEN_SIZE, results)
    # Set fixed window size
    app.resizable(False, False)
    app.geometry(f"{CANVAS_WIDTH + 2 * app.padding}x{CANVAS_HEIGHT + 22 * app.padding}")

    app.mainloop()


if __name__ == "__main__":
    main()
