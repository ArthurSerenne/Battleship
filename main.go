package main

import (
	"battleship/client"
	"battleship/game"
	"battleship/server"
	"battleship/ui"
	"flag"
	"fmt"
	"os"
)

func main() {
	port := flag.String("port", "8080", "Port pour mon serveur")
	target := flag.String("target", "http://localhost:8081", "L'adresse de mon adversaire")
	flag.Parse()

	board := game.NewBoard()
	board.InitRandomShips()

	srv := server.NewBattleshipServer(board)
	go func() {
		if err := srv.Start(*port); err != nil {
			fmt.Printf("Erreur serveur: %v\n", err)
			os.Exit(1)
		}
	}()

	cl := client.NewClient()

	consoleUI := ui.NewConsoleUI(board, cl, *target)
	consoleUI.Run()
}
