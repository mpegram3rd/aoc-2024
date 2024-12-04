package main

// Struct to accumulate data from the input file
type SolutionAccumulator struct {
	data [][]rune
}

// Process and accumulate data from the file
func (acc *SolutionAccumulator) ProcessLine(line string) {
	acc.data = append(acc.data, []rune(line))
}

// Execute the strategies for each solution
func (acc *SolutionAccumulator) Execute() {

	solution1(acc.data)
	solution2(acc.data)
}

// Part 1 solution
func solution1(data [][]rune) {
}

// Part 2 solution
func solution2(data [][]rune) {
}
