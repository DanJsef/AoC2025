package day08

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Distance struct {
	distance    int
	junctionBox [2]*JunctionBox
}

type JunctionBox struct {
	id          int
	x           int
	y           int
	z           int
	circuitId   int
	computedIds map[int]bool
}

func distanceSquared(p1, p2 *JunctionBox) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z

	return dx*dx + dy*dy + dz*dz
}

func Run() {
	file, err := os.Open("./inputs/day08.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	junctionBoxMap := map[int]*JunctionBox{}
	highestId := 0

	circuits := map[int][]*JunctionBox{}

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ",")

		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])

		jb := &JunctionBox{id: highestId, x: x, y: y, z: z, computedIds: map[int]bool{}, circuitId: highestId}

		junctionBoxMap[highestId] = jb
		circuits[highestId] = []*JunctionBox{jb}

		highestId++
	}

	distances := []Distance{}

	for i := 0; i < highestId; i++ {
		for n := 0; n < highestId; n++ {
			if i == n {
				continue
			}

			a := junctionBoxMap[i]

			if _, ok := a.computedIds[n]; ok {
				continue
			}

			b := junctionBoxMap[n]

			distance := distanceSquared(a, b)

			a.computedIds[b.id] = true
			b.computedIds[a.id] = true

			distances = append(distances, Distance{distance: distance, junctionBox: [2]*JunctionBox{a, b}})
		}
	}

	slices.SortFunc(distances, func(a, b Distance) int { return a.distance - b.distance })

	connectionsMade := 0

	for _, d := range distances {
		if connectionsMade == 1000 {
			circuitSizes := []int{}

			for _, c := range circuits {
				circuitSizes = append(circuitSizes, len(c))
			}

			slices.SortFunc(circuitSizes, func(a, b int) int { return b - a })

			fmt.Println(circuitSizes[0] * circuitSizes[1] * circuitSizes[2])
		}

		a, b := d.junctionBox[0], d.junctionBox[1]

		if a.circuitId == b.circuitId {
			connectionsMade++
			continue
		}

		bCircuit := circuits[b.circuitId]
		bCircuitId := b.circuitId

		for _, j := range bCircuit {
			j.circuitId = a.circuitId
			circuits[a.circuitId] = append(circuits[a.circuitId], j)
		}

		delete(circuits, bCircuitId)
		connectionsMade++
		if len(circuits) == 1 {
			fmt.Println(a.x * b.x)
			break
		}
	}
}
