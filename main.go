package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DanJsef/AoC2025/puzzles/day01"
	"github.com/DanJsef/AoC2025/puzzles/day02"
	"github.com/DanJsef/AoC2025/puzzles/day03"
	"github.com/DanJsef/AoC2025/puzzles/day04"
	"github.com/DanJsef/AoC2025/puzzles/day05"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Run day:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	dayIdx, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Invalid day input")
		return
	}

	switch int(dayIdx) {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 3:
		day03.Run()
	case 4:
		day04.Run()
	case 5:
		day05.Run()
	default:
		fmt.Println("Invalid day input")
	}
}
