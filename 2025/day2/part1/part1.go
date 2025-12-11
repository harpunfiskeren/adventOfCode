package part1

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
		current := oneRange[0]
		fmt.Println(oneRange)
		for current <= oneRange[1] {
			str := strconv.Itoa(current)
			if str[:len(str)/2] == str[len(str)/2:] {
				total += current
				fmt.Println(str)
			}
			current++
		}
	}
	fmt.Println(total)
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
