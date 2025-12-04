package data_structs

import "math"

type Position struct {
	X int
	Y int
}

func (p Position) Add(p2 Position) Position {
	return Position{X: p.X + p2.X, Y: p.Y + p2.Y}
}

func (p Position) Sub(p2 Position) Position {
	return Position{X: p.X - p2.X, Y: p.Y - p2.Y}
}

func (p Position) AddWrap(p2 Position, width int, height int) Position {
	newPos := Position{X: (p.X + p2.X) % width, Y: (p.Y + p2.Y) % height}

	if newPos.X < 0 {
		newPos.X += width
	}

	if newPos.Y < 0 {
		newPos.Y += height
	}

	return newPos
}

func (p Position) RotateClockwise() Position {
	return Position{X: -p.Y, Y: p.X}
}

func (p Position) RotateCounterClockwise() Position {
	return Position{X: p.Y, Y: -p.X}
}

func (p Position) IsWithinBounds(width int, height int) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}

func (p Position) ManhattanDistance(p2 Position) int {
	return int(math.Abs(float64(p.X-p2.X)) + math.Abs(float64(p.Y-p2.Y)))
}
