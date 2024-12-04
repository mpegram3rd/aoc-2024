package main

import "fmt"

var matchWord = []rune("XMAS")

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
	var sum = 0
	var y = 0
	for _, row := range data {
		var x = 0
		for _, char := range row {
			if char == 'X' {
				sum += wordSearch(data, x, y)
			}
			x++
		}
		y++
	}
	fmt.Println("Solution 1: ", sum)
}

// Part 2 solution
func solution2(data [][]rune) {
	var sum = 0
	var y = 0
	for _, row := range data {
		var x = 0
		for _, char := range row {
			if char == 'A' {
				sum += checkXmas(data, x, y)
			}
			x++
		}
		y++
	}
	fmt.Println("Solution 2: ", sum)
}

// Solution 1 Logic - Searches the rune from the start point in all directions to see if it matches the matchWord (solution 1)
func wordSearch(data [][]rune, x, y int) int {
	var matches = checkDirection(data, x, y, 1, 0) // Check right
	matches += checkDirection(data, x, y, -1, 0)   // Check left
	matches += checkDirection(data, x, y, 0, 1)    // Check down
	matches += checkDirection(data, x, y, 0, -1)   // Check up
	matches += checkDirection(data, x, y, 1, 1)    // Check down right
	matches += checkDirection(data, x, y, 1, -1)   // Check up right
	matches += checkDirection(data, x, y, -1, 1)   // Check down left
	matches += checkDirection(data, x, y, -1, -1)  // Check up left

	return matches
}

// Solution 1 Logic - Starting from the letter X checks in the given direction for the string XMAS
func checkDirection(data [][]rune, x, y, horiz, vert int) int {
	var curX = x
	var curY = y
	var matching = true
	for wi, char := range matchWord {
		matching = char == data[curY][curX]
		if !matching {
			return 0
		}
		if wi < len(matchWord)-1 {
			curX += horiz
			curY += vert
			if curX < 0 || curY < 0 || curY >= len(data) || curX >= len(data[curY]) {
				return 0
			}
		}

	}
	return 1
}

// Solution 2 Logic - Centered on an "A" checks both diagonals to see if they spell "MAS"
// IE: an "X" formed with MAS, hence X-MAS
func checkXmas(data [][]rune, x, y int) int {
	// no sense checking if x or y are at the edges since it is the center of the cross
	if x < 1 || y < 1 || y > len(data)-2 || x > len(data[y])-2 {
		return 0
	}

	var matches = hasMAS(data, x-1, y-1, 1)
	matches = matches && hasMAS(data, x-1, y+1, -1) // check diagonal from left to right up
	if matches {
		return 1
	}
	return 0
}

// Check whether the diagonal spells the word MAS
func hasMAS(data [][]rune, x int, y int, yDir int) bool {
	var hasM = data[y][x] == 'M'
	var hasS = data[y][x] == 'S'

	// skip the A and check the other letter in the diagonal direction
	x += 2
	y += yDir * 2
	return (hasM && data[y][x] == 'S') || (hasS && data[y][x] == 'M')
}
