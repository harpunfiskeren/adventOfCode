package part2

import (
	"adventOfCode"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var dial = 50
var zeroCount = 0

type Direction string

const (
	left  = "L"
	right = "R"
)

type Instruction struct {
	direction Direction
	steps     int
}

func moveDial(instruction Instruction) {
	if instruction.direction == right {
		handleRightTurn(instruction.steps)
	} else if instruction.direction == left {
		handleLeftTurn(instruction.steps)
	}
}

func handleRightTurn(steps int) {
	dial += steps
	zeroCount += dial / 100
	dial %= 100
}

func handleLeftTurn(steps int) {
	wasZero := dial == 0
	dial -= steps
	landedOnZero := dial%100 == 0

	if landedOnZero {
		zeroCount++
	}

	for dial < 0 {
		dial += 100
		zeroCount++
	}

	if wasZero {
		zeroCount--
	}
}

func Run() {
	file, _ := adventOfCode.OpenInput("2025", "day1", "input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		instruction := parseInputLineToInstruction(line)
		moveDial(instruction)
	}
	fmt.Println("Total times dial was 0: ", zeroCount)
}

func parseInputLineToInstruction(instruction string) Instruction {
	steps, _ := strconv.Atoi(instruction[1:])

	if strings.Contains(instruction, "L") {
		return Instruction{direction: left, steps: steps}
	}
	return Instruction{direction: right, steps: steps}
}
