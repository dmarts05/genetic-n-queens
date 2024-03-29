import os
import platform
import subprocess
import threading
import tkinter as tk
from tkinter import messagebox, ttk
from typing import Tuple


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

        selection_method_label = ttk.Label(self, text="Selection Method:")
        selection_method_label.grid(
            row=len(labels) + 1, column=0, sticky="w", padx=5, pady=5
        )
        selection_methods = ["tournament", "roulette"]
        self.selection_method_combo = ttk.Combobox(self, values=selection_methods)
        self.selection_method_combo.grid(row=len(labels) + 1, column=1, padx=5, pady=5)

        # Show submit button along with loading label
        self.submit_button = ttk.Button(self, text="Submit", command=self.submit)
        self.submit_button.grid(row=len(labels) + 2, column=1, columnspan=1, pady=10)

        self.loading_label = ttk.Label(self, text="Loading...")
        self.loading_label.grid(row=len(labels) + 2, column=0, columnspan=1, pady=10)
        self.loading_label.grid_remove()

    def load_default_values(self):
        default_values = [5, 16, 500, 8, 0.2, 0.5]
        for entry, value in zip(self.entries, default_values):
            entry.delete(0, tk.END)
            entry.insert(0, str(value))

        self.elitism_var.set(False)
        self.selection_method_combo.set("tournament")

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
            selection_method_str = self.selection_method_combo.get()
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
            f"-selectionMethod={selection_method_str}",
        )

        threading.Thread(target=self.process_submission, args=(args,)).start()

    def process_submission(self, args: Tuple[str, ...]) -> None:
        if self.start_golang_implementation(args):
            # Hide loading label before showing results
            self.loading_label.grid_remove()
            self.show_results()

        # Enable submit button
        self.submit_button.config(state="normal")


if __name__ == "__main__":
    app = GeneticNQueensApp()
    app.resizable(False, False)
    app.geometry("300x310")
    app.mainloop()
