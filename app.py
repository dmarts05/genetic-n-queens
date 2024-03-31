from enum import Enum
import os
import platform
import subprocess
import threading
import tkinter as tk
from tkinter import messagebox, ttk
from typing import Tuple


class Implementation(Enum):
    GO = "go"
    DEAP = "deap"


class SelectionMethod(Enum):
    TOURNAMENT = "tournament"
    ROULETTE = "roulette"


class GeneticNQueensApp(tk.Tk):
    def __init__(self):
        super().__init__()
        self.title("Genetic N-Queens")
        self.configure(padx=10, pady=10)
        self.create_widgets()
        self.load_default_values()

    def create_widgets(self):
        labels = [
            "Number of Runs",
            "Population Size",
            "Max Generations",
            "Number of Queens",
            "Mutation Rate",
            "Crossover Rate",
        ]
        self.entries: list[ttk.Entry] = []
        for i, label_text in enumerate(labels):
            label = ttk.Label(self, text=label_text + ":")
            label.grid(row=i, column=0, sticky="w", padx=5, pady=5)

            entry = ttk.Entry(self)
            entry.grid(row=i, column=1, padx=5, pady=5)
            self.entries.append(entry)

        self.elitism_var = tk.BooleanVar()
        elitism_checkbox = ttk.Checkbutton(
            self, text="Elitism", variable=self.elitism_var
        )
        elitism_checkbox.grid(row=len(labels), column=1, columnspan=1, padx=5, pady=5)

        self.selection_method_label = ttk.Label(self, text="Selection Method:")
        self.selection_method_label.grid(
            row=len(labels) + 1, column=0, sticky="w", padx=5, pady=5
        )
        self.selection_methods = [
            SelectionMethod.TOURNAMENT.value,
            SelectionMethod.ROULETTE.value,
        ]
        self.selection_method_combo = ttk.Combobox(self, values=self.selection_methods)
        self.selection_method_combo.grid(row=len(labels) + 1, column=1, padx=5, pady=5)
        self.selection_method_combo.bind(
            "<<ComboboxSelected>>",
            self.toggle_tournament_size,  # type: ignore
        )

        self.tournament_size_label = ttk.Label(self, text="Tournament Size:")
        self.tournament_size_entry = ttk.Entry(self)
        self.tournament_size_label.grid(
            row=len(labels) + 2, column=0, sticky="w", padx=5, pady=5
        )
        self.tournament_size_entry.grid(row=len(labels) + 2, column=1, padx=5, pady=5)
        self.hide_tournament_size_entry()

        implementations_label = ttk.Label(self, text="Implementation:")
        implementations_label.grid(
            row=len(labels) + 3, column=0, sticky="w", padx=5, pady=5
        )
        implementations = [Implementation.GO.value, Implementation.DEAP.value]
        self.implementations_combo = ttk.Combobox(self, values=implementations)
        self.implementations_combo.grid(row=len(labels) + 3, column=1, padx=5, pady=5)

        # Show submit button along with loading label
        self.submit_button = ttk.Button(self, text="Submit", command=self.submit)
        self.submit_button.grid(row=len(labels) + 4, column=1, columnspan=1, pady=10)

        self.loading_label = ttk.Label(self, text="Loading...")
        self.loading_label.grid(row=len(labels) + 4, column=0, columnspan=1, pady=10)
        self.loading_label.grid_remove()

    def load_default_values(self):
        num_runs = 12
        population_size = 300
        max_generations = 3000
        num_queens = 29
        mutation_rate = 0.2
        crossover_rate = 0.5
        default_values = [
            num_runs,
            population_size,
            max_generations,
            num_queens,
            mutation_rate,
            crossover_rate,
        ]
        for entry, value in zip(self.entries, default_values):
            entry.delete(0, tk.END)
            entry.insert(0, str(value))

        self.elitism_var.set(False)
        self.selection_method_combo.set(SelectionMethod.TOURNAMENT.value)
        self.tournament_size_entry.delete(0, tk.END)
        self.tournament_size_entry.insert(0, "3")
        self.toggle_tournament_size()  # type: ignore
        self.implementations_combo.set(Implementation.GO.value)

    def toggle_tournament_size(self, event=None):  # type: ignore
        selection_method = self.selection_method_combo.get()
        if selection_method == SelectionMethod.TOURNAMENT.value:
            self.show_tournament_size_entry()
        else:
            self.hide_tournament_size_entry()

    def show_tournament_size_entry(self):
        self.tournament_size_label.grid()
        self.tournament_size_entry.grid()

    def hide_tournament_size_entry(self):
        self.tournament_size_label.grid_remove()
        self.tournament_size_entry.grid_remove()

    def start_deap_implementation(self, args: Tuple[str, ...]) -> bool:
        try:
            subprocess.run(["python", "deap-implementation.py", *args], check=True)
        except subprocess.CalledProcessError:
            messagebox.showerror("Error", "Failed to start the DEAP app.")

    def start_golang_implementation(self, args: Tuple[str, ...]) -> bool:
        def get_executable_path() -> str:
            system = platform.system().lower()
            executables = {
                "windows": "genetic-n-queens-windows-amd64.exe",
                "linux": "genetic-n-queens-linux-amd64",
                "darwin": "genetic-n-queens-darwin-amd64",
            }
            if system not in executables:
                messagebox.showerror("Error", "Unsupported OS")  # type: ignore
                return ""
            return os.path.join("go-implementation", "bin", executables[system])

        def build_executable() -> bool:
            ok = True

            os.chdir("go-implementation")
            try:
                subprocess.run(["make", "build/prod"], check=True)
            except subprocess.CalledProcessError:
                messagebox.showerror(  # type: ignore
                    "Error",
                    "Failed to build the binary. Ensure you have Make and Go installed or download the binary from the releases page.",
                )
                ok = False
            finally:
                os.chdir("..")

            return ok

        executable = get_executable_path()
        if not executable or not os.path.exists(executable) and not build_executable():
            return False

        try:
            subprocess.run([executable, *args], check=True)
        except subprocess.CalledProcessError:
            messagebox.showerror("Error", "Failed to start the Golang app.")  # type: ignore
            return False

        return True

    def show_results(self) -> None:
        try:
            subprocess.run(["python", "show_results.py", "results.json"])
        except subprocess.CalledProcessError:
            messagebox.showerror("Error", "Failed to show the results.")  # type: ignore

    def submit(self) -> None:
        try:
            num_runs = int(self.entries[0].get())
            population_size = int(self.entries[1].get())
            max_generations = int(self.entries[2].get())
            num_queens = int(self.entries[3].get())
            mutation_rate = float(self.entries[4].get())
            crossover_rate = float(self.entries[5].get())
            elitism = self.elitism_var.get()
            selection_method = SelectionMethod(self.selection_method_combo.get())
            tournament_size = (
                int(self.tournament_size_entry.get())
                if self.selection_method_combo.get() == SelectionMethod.TOURNAMENT.value
                else None
            )
            implementation = Implementation(self.implementations_combo.get())
        except ValueError:
            messagebox.showerror("Error", "Invalid input")  # type: ignore
            return

        # Disable submit button and show loading label
        self.submit_button.config(state="disabled")
        self.loading_label.grid()

        args = (
            f"-numRuns={num_runs}",
            f"-populationSize={population_size}",
            f"-maxGenerations={max_generations}",
            f"-numQueens={num_queens}",
            f"-mutationRate={mutation_rate}",
            f"-crossOverRate={crossover_rate}",
            f"-elitism={elitism}",
            f"-selectionMethod={selection_method.value}",
            f"-tournamentSize={tournament_size}" if tournament_size is not None else "",
        )

        threading.Thread(
            target=self.process_submission, implementation=implementation, args=(args,)
        ).start()

    def process_submission(
        self, implementation: Implementation, args: Tuple[str, ...]
    ) -> None:
        match implementation:
            case Implementation.DEAP:
                ok = self.start_deap_implementation(args)
            case Implementation.GO:
                ok = self.start_golang_implementation(args)

        if ok:
            # Hide loading label before showing results
            self.loading_label.grid_remove()
            self.show_results()

        # Enable submit button
        self.submit_button.config(state="normal")


if __name__ == "__main__":
    app = GeneticNQueensApp()
    app.minsize(300, 370)
    app.mainloop()
