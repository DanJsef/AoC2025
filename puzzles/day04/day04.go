package day04

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DanJsef/AoC2025/internal/data_structs"
)

var directions = []data_structs.Position{
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
	{X: 1, Y: 1},
	{X: -1, Y: 1},
	{X: -1, Y: -1},
	{X: 1, Y: -1},
}

type Plan struct {
	value           [][]rune
	sizeX           int
	sizeY           int
	accessibleRolls int
	toRemove        []data_structs.Position
}

func (p *Plan) print() {
	for _, v := range p.value {
		fmt.Println(string(v))
	}
}

func (p *Plan) checkAdjacent(pos data_structs.Position) {
	count := 0

	for _, direction := range directions {
		newPos := pos.Add(direction)
		if !newPos.IsWithinBounds(p.sizeX, p.sizeY) {
			continue
		}

		if p.value[newPos.Y][newPos.X] == '@' {
			count++
		}
	}

	if count < 4 {
		p.accessibleRolls++
		p.toRemove = append(p.toRemove, pos)
	}
}

func (p *Plan) countAccessibleRolls() {
	for y := range p.value {
		for x := range p.value[y] {
			if p.value[y][x] == '.' {
				continue
			}

			p.checkAdjacent(data_structs.Position{X: x, Y: y})
		}
	}
}

func (p *Plan) removeRolls() {
	for _, pos := range p.toRemove {
		p.value[pos.Y][pos.X] = '.'
	}
	p.toRemove = []data_structs.Position{}
}

func Run() {
	file, err := os.Open("./inputs/day04.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	plan := Plan{sizeX: 0, sizeY: 0}

	for scanner.Scan() {
		line := scanner.Text()

		plan.sizeX = len(line)
		plan.value = append(plan.value, []rune(line))
		plan.sizeY++
	}

	plan.countAccessibleRolls()

	fmt.Println(plan.accessibleRolls)

	for len(plan.toRemove) > 0 {
		plan.removeRolls()
		plan.countAccessibleRolls()
	}

	fmt.Println(plan.accessibleRolls)
}
