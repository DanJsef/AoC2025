package day06

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	numbers   []int
	operation string
	result    int
}

func Run() {
	file, err := os.Open("./inputs/day06.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var worksheet []Problem

	worksheet2 := [][]rune{}
	var operatorIndexes []int

	numMode := true

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)

		if len(worksheet) == 0 {
			worksheet = make([]Problem, len(split))
		}

		worksheet2 = append(worksheet2, []rune(line))

		for i, val := range split {
			num, err := strconv.Atoi(val)
			if err != nil {
				numMode = false
				break
			}

			worksheet[i].numbers = append(worksheet[i].numbers, num)
		}

		if numMode {
			continue
		}

		for i, val := range split {
			switch val {
			case "+":
				worksheet[i].result = 0
			case "*":
				worksheet[i].result = 1
			}

			worksheet[i].operation = val
		}

		operatorIndexes = make([]int, len(split))

		idxPointer := 0
		for i, rune := range line {
			if rune != ' ' {
				operatorIndexes[idxPointer] = i
				idxPointer++
			}
		}
	}

	grandTotal := 0

	for _, problem := range worksheet {
		for _, num := range problem.numbers {
			switch problem.operation {
			case "*":
				problem.result *= num

			case "+":
				problem.result += num
			}
		}
		grandTotal += problem.result
	}

	fmt.Println(grandTotal)

	grandTotal2 := 0

	for i, opIdx := range operatorIndexes {
		operator := worksheet2[len(worksheet2)-1][opIdx]

		rightBound := len(worksheet2[0]) - 1
		if i != len(operatorIndexes)-1 {
			rightBound = operatorIndexes[i+1] - 2
		}

		total := 0
		if operator == '*' {
			total = 1
		}

		for ; rightBound >= opIdx; rightBound-- {
			numberRunes := []rune{}

			for upBound := 0; upBound < len(worksheet2)-1; upBound++ {
				if worksheet2[upBound][rightBound] == ' ' {
					continue
				}

				numberRunes = append(numberRunes, worksheet2[upBound][rightBound])
			}

			number, _ := strconv.Atoi(string(numberRunes))

			switch operator {
			case '*':
				total *= number

			case '+':
				total += number
			}
		}
		grandTotal2 += total
	}

	fmt.Println(grandTotal2)
}
