package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Safe struct {
	lockValue      int
	zeroCount      int
	zeroClickCount int
}

func (s *Safe) rotateLeft(byValue int) {
	old := s.lockValue
	s.lockValue -= byValue

	if s.lockValue >= 0 {
		return
	}

	s.lockValue = 99 + s.lockValue + 1

	if s.lockValue != 0 && old != 0 {
		fmt.Println("Increase count")
		s.zeroClickCount++
	}
}

func (s *Safe) rotateRight(byValue int) {
	old := s.lockValue
	s.lockValue += byValue

	if s.lockValue <= 99 {
		return
	}

	s.lockValue = s.lockValue - 99 - 1

	if s.lockValue != 0 && old != 0 {
		fmt.Println("Increase count")
		s.zeroClickCount++
	}
}

func (s *Safe) checkLockState() {
	if s.lockValue != 0 {
		return
	}

	s.zeroCount++
	s.zeroClickCount++
}

func Run() {
	fmt.Println("hello day 01")

	file, err := os.Open("./inputs/day01.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	safe := Safe{lockValue: 50, zeroCount: 0}

	for reader.Scan() {
		line := reader.Text()
		direction := line[:1]
		amount, _ := strconv.Atoi(line[1:])

		safe.zeroClickCount += amount / 100

		amount = amount % 100

		if direction == "R" {
			safe.rotateRight(amount)
		}
		if direction == "L" {
			safe.rotateLeft(amount)
		}

		fmt.Println(safe.lockValue)
		fmt.Println()
		safe.checkLockState()
	}

	fmt.Println(strconv.Itoa(safe.zeroCount))
	fmt.Println(strconv.Itoa(safe.zeroClickCount))
}
