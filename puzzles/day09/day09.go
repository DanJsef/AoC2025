package day09

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/DanJsef/AoC2025/internal/data_structs"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

type Segment struct {
	edges      []data_structs.Position
	isVertical bool
}

func (s *Segment) prepare() {
	s.isVertical = s.edges[0].X == s.edges[1].X
}

type Candidate struct {
	edges [2]data_structs.Position
	size  int
}

func (c *Candidate) getSegments() [4]Segment {
	p1 := c.edges[0]
	p2 := c.edges[1]
	p3 := data_structs.Position{X: p1.X, Y: p2.Y}
	p4 := data_structs.Position{X: p2.X, Y: p1.Y}

	return [4]Segment{
		{edges: []data_structs.Position{p1, p3}},
		{edges: []data_structs.Position{p3, p2}},
		{edges: []data_structs.Position{p2, p4}},
		{edges: []data_structs.Position{p4, p1}},
	}
}

func IsSegmentEnclosed(segments []Segment, testSegment Segment) bool {

	start := testSegment.edges[0]
	end := testSegment.edges[1]

	if !isPointInsideOrOnBoundary(segments, start) {
		return false
	}
	if !isPointInsideOrOnBoundary(segments, end) {
		return false
	}

	for _, boundarySeg := range segments {
		b1 := boundarySeg.edges[0]
		b2 := boundarySeg.edges[1]

		if properSegmentIntersection(start, end, b1, b2) {
			return false
		}
	}

	return true
}

func isPointInsideOrOnBoundary(segments []Segment, point data_structs.Position) bool {
	for _, seg := range segments {
		if isPointOnSegment(seg.edges[0], seg.edges[1], point) {
			return true
		}
	}

	intersections := 0

	for _, seg := range segments {
		p1 := seg.edges[0]
		p2 := seg.edges[1]

		if rayIntersectsSegment(point, p1, p2) {
			intersections++
		}
	}

	return intersections%2 == 1
}

func isPointOnSegment(p1, p2, point data_structs.Position) bool {
	cross := (point.Y-p1.Y)*(p2.X-p1.X) - (point.X-p1.X)*(p2.Y-p1.Y)

	if math.Abs(float64(cross)) > 1e-9 {
		return false
	}

	minX := min(p1.X, p2.X)
	maxX := max(p1.X, p2.X)
	minY := min(p1.Y, p2.Y)
	maxY := max(p1.Y, p2.Y)

	return point.X >= minX && point.X <= maxX && point.Y >= minY && point.Y <= maxY
}

func rayIntersectsSegment(point, p1, p2 data_structs.Position) bool {
	if p1.Y == p2.Y {
		return false
	}

	minY := min(p1.Y, p2.Y)
	maxY := max(p1.Y, p2.Y)

	if point.Y < minY || point.Y >= maxY {
		return false
	}

	t := float64(point.Y-p1.Y) / float64(p2.Y-p1.Y)
	xIntersect := float64(p1.X) + t*float64(p2.X-p1.X)

	return xIntersect > float64(point.X)
}

func properSegmentIntersection(a1, a2, b1, b2 data_structs.Position) bool {
	d1 := direction(b1, b2, a1)
	d2 := direction(b1, b2, a2)
	d3 := direction(a1, a2, b1)
	d4 := direction(a1, a2, b2)

	if ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) &&
		((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0)) {
		return true
	}

	return false
}

func direction(p1, p2, p3 data_structs.Position) float64 {
	return float64((p3.X-p1.X)*(p2.Y-p1.Y) - (p2.X-p1.X)*(p3.Y-p1.Y))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Run() {
	file, err := os.Open("./inputs/day09.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	positions := []data_structs.Position{}

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ",")

		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])

		positions = append(positions, data_structs.Position{X: x, Y: y})
	}

	largest := 0

	candidates := []Candidate{}

	for i, p := range positions {
		for j, q := range positions {
			if i == j {
				continue
			}

			size := (abs(p.X-q.X) + 1) * (abs(p.Y-q.Y) + 1)

			candidates = append(candidates, Candidate{edges: [2]data_structs.Position{p, q}, size: size})

			if size > largest {
				largest = size
			}
		}
	}

	fmt.Println(largest)

	slices.SortFunc(candidates, func(a, b Candidate) int { return b.size - a.size })

	segments := make([]Segment, len(positions))

	for i, p := range positions {
		next := (i + 1) % len(positions)

		segments[i] = Segment{edges: []data_structs.Position{p, positions[next]}}
		segments[i].prepare()
	}

	slices.SortFunc(segments, func(a, b Segment) int { return a.edges[0].Y - b.edges[0].Y })

	for i, candidate := range candidates {
		testSegments := candidate.getSegments()

		miss := false
		for _, ts := range testSegments {
			ts.prepare()

			if !IsSegmentEnclosed(segments, ts) {
				miss = true
				break
			}
		}

		if !miss {
			fmt.Println(candidates[i].size)
			break
		}
	}
}
