package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// RegEx "constants"
var parserRe = regexp.MustCompile(`do\(\)|don't\(\)|mul\([0-9]+,[0-9]+\)`)
var valuesRe = regexp.MustCompile(`[0-9]+`)
var multRe = regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)

// Struct to accumulate data from the input file
type SolutionAccumulator struct {
	sb strings.Builder
}

// Process and accumulate data from the file
func (acc *SolutionAccumulator) ProcessLine(line string) {
	acc.sb.WriteString(line)
}

// Execute the strategies for each solution
func (acc *SolutionAccumulator) Execute() {
	var memoryDump = acc.sb.String()

	solution1(memoryDump)
	solution2(memoryDump)
}

// Part 1 solution
func solution1(memoryDump string) {
	var operations = multRe.FindAllString(memoryDump, -1)
	var sum = 0
	for _, operation := range operations {
		sum += multiply(operation)
	}
	fmt.Println("Solution 1: ", sum)
}

// Part 2 solution
func solution2(memoryDump string) {
	var operations = parserRe.FindAllString(memoryDump, -1)
	var sum = 0
	var process = true
	for _, operation := range operations {
		if strings.HasPrefix(operation, "do(") {
			process = true
		}
		if strings.HasPrefix(operation, "don't(") {
			process = false
		}
		if process && strings.HasPrefix(operation, "mul(") {
			sum += multiply(operation)
		}
	}
	fmt.Println("Solution 1: ", sum)
}

// process the multiply command
func multiply(operation string) int {
	var params = valuesRe.FindAllString(operation, -1)
	if len(params) == 2 {
		var param1, _ = strconv.Atoi(params[0])
		var param2, _ = strconv.Atoi(params[1])
		return param1 * param2
	}
	return 0
}
