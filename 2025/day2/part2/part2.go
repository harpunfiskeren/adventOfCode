package part2

import (
	"adventOfCode"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	file, _ := adventOfCode.OpenInput("2025", "day2", "input")
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	ranges := parseInputToRanges(scanner.Text())
	var total int

	for _, oneRange := range ranges {
		currentNumberInRange := oneRange[0]

		for currentNumberInRange <= oneRange[1] {
			str := strconv.Itoa(currentNumberInRange)

			if isInvalid(str) {
				total += currentNumberInRange
			}
			currentNumberInRange++
		}
	}
	fmt.Println("Result part 2: ", total)
}

func isInvalid(str string) bool {
	for i := 1; i <= len(str)/2; i++ {
		sequence := str[:i]
		expectedStr := ""
		for i := 0; i < len(str)/len(sequence); i++ {
			expectedStr += sequence
		}

		if str == expectedStr {
			return true
		}
	}
	return false
}

func parseInputToRanges(s string) [][]int {

	var ranges [][]int
	split := strings.Split(s, ",")
	for i := 0; i < len(split); i++ {
		pair := strings.Split(split[i], "-")
		numPairs := make([]int, len(pair))
		for num := range numPairs {
			numPairs[num], _ = strconv.Atoi(pair[num])
		}
		ranges = append(ranges, numPairs)
	}
	return ranges
}
