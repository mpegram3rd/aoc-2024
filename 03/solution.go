package main

import "strings"

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

}

// Part 2 solution
func solution2(memoryDump string) {
}
