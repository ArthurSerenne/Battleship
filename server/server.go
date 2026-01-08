package server

import (
	"battleship/game"
	"encoding/json"
	"fmt"
	"net/http"
)

type BattleshipServer struct {
	Board *game.Board
}

func NewBattleshipServer(b *game.Board) *BattleshipServer {
	return &BattleshipServer{Board: b}
}

func (s *BattleshipServer) Start(port string) error {
	http.HandleFunc("/hit", s.handleHit)
	http.HandleFunc("/boats", s.handleBoats)
	http.HandleFunc("/board", s.handleBoard)

	fmt.Printf("Serveur démarré sur le port %s...\n", port)
	return http.ListenAndServe(":"+port, nil)
}

type FireRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type FireResponse struct {
	Result string `json:"result"`
}

func (s *BattleshipServer) handleHit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	if s.Board.CountShipsAlive() == 0 {
		http.Error(w, "Le joueur a déjà perdu", http.StatusForbidden)
		return
	}

	var req FireRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Mauvais format JSON", http.StatusBadRequest)
		return
	}

	result := s.Board.ReceiveHit(req.X, req.Y)

	resp := FireResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *BattleshipServer) handleBoats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	count := s.Board.CountShipsAlive()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"boats_alive": count})
}

func (s *BattleshipServer) handleBoard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.Board)
}
