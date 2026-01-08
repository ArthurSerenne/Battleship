package game

import (
	"math/rand"
	"time"
)

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) InitRandomShips() {
	rand.Seed(time.Now().UnixNano())
	sizes := []int{5, 4, 3, 3, 2}

	for _, size := range sizes {
		placed := false
		for !placed {
			horizontal := rand.Intn(2) == 0
			var x, y int
			if horizontal {
				x = rand.Intn(Size - size + 1)
				y = rand.Intn(Size)
			} else {
				x = rand.Intn(Size)
				y = rand.Intn(Size - size + 1)
			}

			if b.canPlace(x, y, size, horizontal) {
				b.place(x, y, size, horizontal)
				placed = true
			}
		}
	}
}

func (b *Board) canPlace(x, y, size int, horizontal bool) bool {
	for i := 0; i < size; i++ {
		cx, cy := x, y
		if horizontal {
			cx += i
		} else {
			cy += i
		}

		if b.Grid[cy][cx] == CellShip {
			return false
		}
	}
	return true
}

func (b *Board) place(x, y, size int, horizontal bool) {
	newShip := &Ship{Coords: []Coordinate{}}
	for i := 0; i < size; i++ {
		cx, cy := x, y
		if horizontal {
			cx += i
		} else {
			cy += i
		}
		b.Grid[cy][cx] = CellShip
		newShip.Coords = append(newShip.Coords, Coordinate{X: cx, Y: cy})
	}
	b.Ships = append(b.Ships, newShip)
}

func (s *Ship) IsAlive(grid [Size][Size]int) bool {
	for _, c := range s.Coords {
		if grid[c.Y][c.X] != CellHit {
			return true
		}
	}
	return false
}

func (b *Board) CountShipsAlive() int {
	count := 0
	for _, s := range b.Ships {
		if s.IsAlive(b.Grid) {
			count++
		}
	}
	return count
}

func (b *Board) ReceiveHit(x, y int) string {
	if x < 0 || x >= Size || y < 0 || y >= Size {
		return "hors_limite"
	}

	val := b.Grid[y][x]

	if val == CellHit || val == CellMiss {
		return "deja_joue"
	}

	if val == CellShip {
		b.Grid[y][x] = CellHit
		return "touché"
	}

	b.Grid[y][x] = CellMiss
	return "plouf"
}
