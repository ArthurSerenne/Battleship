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
	TargetURL    string
	TrackingGrid [game.Size][game.Size]int
}

func NewConsoleUI(b *game.Board, c *client.Client, target string) *ConsoleUI {
	return &ConsoleUI{
		Board:     b,
		Client:    c,
		TargetURL: target,
	}
}

func (ui *ConsoleUI) Run() {
	for {
		ui.ClearScreen()
		ui.PrintState()

		fmt.Println("\nActions:")
		fmt.Println("- Entrez 'x y' pour tirer (ex: '3 4')")
		fmt.Println("- Entrez 'r' pour rafraîchir l'écran")
		fmt.Println("- Entrez 'q' pour quitter")
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

			_, err := fmt.Sscanf(input, "%d", &x)
			if err == nil {
				fmt.Print("Y: ")
				fmt.Scanln(&y)
				ui.fire(x, y)
			}
		}
	}
}

func (ui *ConsoleUI) fire(x, y int) {
	res, err := ui.Client.Fire(ui.TargetURL, x, y)
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
	} else {
		fmt.Printf("Résultat du tir en %d,%d : %s\n", x, y, res)
		if res == "touché" || res == "coulé" {
			ui.TrackingGrid[y][x] = game.CellHit
		} else {
			ui.TrackingGrid[y][x] = game.CellMiss
		}
	}
	fmt.Println("Appuyez sur Entrée pour continuer...")
	var ignore string
	fmt.Scanln(&ignore)
}

func (ui *ConsoleUI) PrintState() {
	fmt.Println("=== BATAILLE NAVALE (Console) ===")

	fmt.Println("\n RADAR (Cible)              MA FLOTTE")
	fmt.Println("   0 1 2 3 4 5 6 7 8 9        0 1 2 3 4 5 6 7 8 9")

	for y := 0; y < game.Size; y++ {
		fmt.Printf("%d ", y)
		for x := 0; x < game.Size; x++ {
			val := ui.TrackingGrid[y][x]
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
