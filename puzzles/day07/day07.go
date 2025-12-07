package day07

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DanJsef/AoC2025/internal/data_structs"
)

type Manifold struct {
	diagram       [][]rune
	currentLevel  int
	splitCount    int
	timelineCount int
	cache         map[[2]int]int
}

func (m *Manifold) step() {
	for x := 0; x < len(m.diagram[0]); x++ {
		if m.diagram[m.currentLevel][x] != '|' && m.diagram[m.currentLevel][x] != 'S' {
			continue
		}

		if m.diagram[m.currentLevel+1][x] != '^' {
			m.diagram[m.currentLevel+1][x] = '|'
			continue
		}

		m.diagram[m.currentLevel+1][x+1], m.diagram[m.currentLevel+1][x-1] = '|', '|'
		m.splitCount++
	}
	m.currentLevel++
}

func (m *Manifold) countTimelines() {
	start := data_structs.Position{X: 0, Y: 0}

	for m.diagram[start.Y][start.X] != 'S' {
		start.X++
	}

	m.timelineCount = m.dfs(start)
}

func (m *Manifold) dfs(pos data_structs.Position) int {
	if cachedCount, ok := m.cache[[2]int{pos.Y, pos.X}]; ok {
		return cachedCount
	}

	if pos.Y == len(m.diagram)-1 {
		return 1
	}

	timelineCount := 0

	if m.diagram[pos.Y+1][pos.X] != '^' {
		timelineCount += m.dfs(data_structs.Position{Y: pos.Y + 1, X: pos.X})
	}

	if m.diagram[pos.Y+1][pos.X] == '^' {
		timelineCount += m.dfs(data_structs.Position{X: pos.X - 1, Y: pos.Y + 1})
		timelineCount += m.dfs(data_structs.Position{X: pos.X + 1, Y: pos.Y + 1})
	}

	m.cache[[2]int{pos.Y, pos.X}] = timelineCount

	return timelineCount
}

func Run() {
	file, err := os.Open("./inputs/day07.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	manifold := Manifold{cache: make(map[[2]int]int)}

	for scanner.Scan() {
		line := scanner.Text()

		manifold.diagram = append(manifold.diagram, []rune(line))
	}

	for manifold.currentLevel < len(manifold.diagram)-1 {
		manifold.step()
	}

	fmt.Println()
	fmt.Println(manifold.splitCount)

	manifold.countTimelines()
	fmt.Println(manifold.timelineCount)
}
