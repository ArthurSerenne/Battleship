package game

const Size = 10

const (
	CellEmpty = 0
	CellMiss  = 1
	CellHit   = 2
	CellShip  = 9
)

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Board struct {
	Grid [Size][Size]int `json:"grid"`
}
