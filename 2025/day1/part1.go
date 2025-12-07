package main

import (
	"adventOfCode"
	"bufio"
	"log"
	"strconv"
	"strings"
)

var dial = 50
var zeroCount = 0

func main() {
	file, _ := adventOfCode.OpenInput("2025", "day1", "input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		instruction := parseInputLineToInstruction(line)
		moveDial(instruction)
	}
	log.Println("Total times dial was 0: ", zeroCount)
}

type Direction string

const (
	left  = "L"
	right = "R"
)

func moveDial(instruction Instruction) {
	if instruction.direction == left {
		dial = (dial - (instruction.steps % 100) + 100) % 100
	}
	if instruction.direction == right {
		dial = (dial + instruction.steps) % 100
	}

	if dial == 0 {
		zeroCount++
	}
}

type Instruction struct {
	direction Direction
	steps     int
}

func parseInputLineToInstruction(instruction string) Instruction {
	steps, _ := strconv.Atoi(instruction[1:])

	if strings.Contains(instruction, "L") {
		return Instruction{direction: left, steps: steps}
	}
	return Instruction{direction: right, steps: steps}
}
