package day05

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkRange(id int, ranges [][2]int) int {
	for _, r := range ranges {
		if r[0] <= id && id <= r[1] {
			return 1
		}
	}
	return 0
}

func countValidIds(ranges [][2]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	count := 0
	lastRange := [2]int{0, 0}

	for _, r := range ranges {
		if r[0] <= lastRange[1] {
			r[0] = lastRange[1] + 1

			if r[1] < r[0] {
				continue
			}
		}

		count += r[1] - r[0] + 1
		lastRange = r
	}

	return count
}

func Run() {
	file, err := os.Open("./inputs/day05.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ranges := [][2]int{}

	ids := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		split := strings.Split(line, "-")

		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])

		ranges = append(ranges, [2]int{left, right})
	}

	for scanner.Scan() {
		line := scanner.Text()

		id, _ := strconv.Atoi(line)

		ids = append(ids, id)
	}

	fresh := 0

	for _, id := range ids {
		fresh += checkRange(id, ranges)
	}

	fmt.Println(fresh)
	fmt.Println(countValidIds(ranges))
}
