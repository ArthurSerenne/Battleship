package ui

import (
	"battleship/client"
	"battleship/game"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type ConsoleUI struct {
	Board        *game.Board
	Client       *client.Client
	Opponents    []string
	TargetURL    string
	TrackingGrid map[string][game.Size][game.Size]int
}

func NewConsoleUI(b *game.Board, c *client.Client, targets []string) *ConsoleUI {
	ui := &ConsoleUI{
		Board:        b,
		Client:       c,
		Opponents:    targets,
		TrackingGrid: make(map[string][game.Size][game.Size]int),
	}
	if len(targets) > 0 {
		ui.TargetURL = targets[0]
	}
	return ui
}

func (ui *ConsoleUI) Run() {
	for {
		ui.ClearScreen()
		ui.PrintState()

		fmt.Printf("Cible actuelle: %s\n", ui.TargetURL)
		fmt.Println("\nActions:")
		fmt.Println("- 'x y' pour tirer (ex: '3 4')")
		fmt.Println("- 'r' pour rafraîchir")
		fmt.Println("- 'q' pour quitter")
		fmt.Print("> ")

		var input string
		fmt.Scanln(&input)

		if input == "q" {
			fmt.Println("Au revoir!")
			os.Exit(0)
		} else if input == "r" {
			continue
		} else {
			var x, y int
			// Petit hack pour relire le premier nombre si c'est un tir
			n, _ := fmt.Sscanf(input, "%d", &x)
			if n > 0 {
				fmt.Print("Y: ")
				fmt.Scanln(&y)
				ui.fire(x, y)
			}
		}
	}
}

func (ui *ConsoleUI) fire(x, y int) {
	if ui.TargetURL == "" {
		fmt.Println("Aucune cible sélectionnée")
		return
	}

	res, err := ui.Client.Fire(ui.TargetURL, x, y)
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
	} else {
		fmt.Printf("Résultat du tir en %d,%d : %s\n", x, y, res)
		grid := ui.TrackingGrid[ui.TargetURL]
		if res == "touché" || res == "coulé" {
			grid[y][x] = game.CellHit
		} else {
			grid[y][x] = game.CellMiss
		}
		ui.TrackingGrid[ui.TargetURL] = grid
	}
	fmt.Println("Appuyez sur Entrée pour continuer...")
	var ignore string
	fmt.Scanln(&ignore)
}

func (ui *ConsoleUI) PrintState() {
	fmt.Println("=== BATAILLE NAVALE (Console) ===")

	fmt.Printf("\n RADAR (%s)              MA FLOTTE\n", ui.TargetURL)
	fmt.Println("   0 1 2 3 4 5 6 7 8 9        0 1 2 3 4 5 6 7 8 9")

	currentGrid := ui.TrackingGrid[ui.TargetURL]

	for y := 0; y < game.Size; y++ {
		fmt.Printf("%d ", y)
		for x := 0; x < game.Size; x++ {
			val := currentGrid[y][x]
			symbol := "~"
			if val == game.CellHit {
				symbol = "X"
			}
			if val == game.CellMiss {
				symbol = "O"
			}
			fmt.Printf(" %s", symbol)
		}

		fmt.Print("      ")

		fmt.Printf("%d ", y)
		for x := 0; x < game.Size; x++ {
			val := ui.Board.Grid[y][x]
			symbol := "~"
			if val == game.CellShip {
				symbol = "B"
			}
			if val == game.CellHit {
				symbol = "X"
			}
			if val == game.CellMiss {
				symbol = "O"
			}
			fmt.Printf(" %s", symbol)
		}
		fmt.Println()
	}
}

func (ui *ConsoleUI) ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
