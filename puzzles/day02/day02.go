package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countDigits(n int) int {
	tmp := n
	digits := 0
	for tmp > 0 {
		tmp /= 10
		digits++
	}
	return digits
}

func hasRepeatingPattern(s string) bool {
	n := len(s)
	if n <= 1 {
		return false
	}

	lps := make([]int, n)
	j := 0
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = lps[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		lps[i] = j
	}
	len := lps[n-1]
	return len > 0 && n%(n-len) == 0
}

func createPattern(n int) int {
	digits := countDigits(n)

	result := n
	for range digits {
		result *= 10
	}
	return result + n
}

func Run() {

	file, err := os.Open("./inputs/day02.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	reader.Scan()
	line := reader.Text()

	ranges := strings.Split(line, ",")

	matchSum := 0

	for i := 1; i < 99999; i++ {
		testValue := createPattern(i)

		for _, r := range ranges {
			parts := strings.Split(r, "-")

			left, _ := strconv.Atoi(parts[0])
			right, _ := strconv.Atoi(parts[1])

			if left <= testValue && testValue <= right {
				matchSum += testValue
			}

		}
	}

	matchSum2 := 0

	for _, r := range ranges {
		parts := strings.Split(r, "-")

		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		for i := left; i <= right; i++ {

			if hasRepeatingPattern(strconv.Itoa(i)) && left <= i && i <= right {
				matchSum2 += i
			}

		}
	}

	fmt.Println()
	fmt.Println(matchSum)
	fmt.Println(matchSum2)
}
