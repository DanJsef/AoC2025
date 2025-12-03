package day03

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func recursiveJoltage(bank []int, depth int) int {
	left := 0
	leftIdx := 0
	l := len(bank)
	for i, battery := range bank {
		if i == l-depth {
			break
		}

		if battery > left {
			leftIdx = i
			left = battery
		}
	}

	if depth == 0 {
		return left
	}

	return left*int(math.Pow10(depth)) + recursiveJoltage(bank[leftIdx+1:], depth-1)
}

func Run() {
	file, err := os.Open("./inputs/day03.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	banks := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		bankSplit := strings.Split(line, "")

		bank := make([]int, len(bankSplit))

		for i, ch := range bankSplit {
			num, err := strconv.Atoi(ch)
			if err != nil {
				fmt.Printf("Error converting '%s': %v\n", ch, err)
				continue
			}
			bank[i] = num
		}

		banks = append(banks, bank)
	}

	sum := 0
	sum2 := 0

	for _, bank := range banks {
		sum += recursiveJoltage(bank, 1)
		sum2 += recursiveJoltage(bank, 11)
	}

	fmt.Println()
	fmt.Println(sum)
	fmt.Println(sum2)
}
