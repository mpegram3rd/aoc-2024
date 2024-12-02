package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Struct to accumulate data from the input file
type SolutionAccumulator struct {
	rawData [][]int
}

// Process and accumulate data from the file
func (acc *SolutionAccumulator) ProcessLine(line string) {
	parts := strings.Split(line, " ")
	var vals []int
	for _, val := range parts {
		intVal, _ := strconv.Atoi(val)
		vals = append(vals, intVal)
	}
	acc.rawData = append(acc.rawData, vals)
}

// Execute the strategies for each solution
func (acc *SolutionAccumulator) Execute() {
	solution1(acc)
	solution2(acc)
}

// Part 1 solution
func solution1(acc *SolutionAccumulator) {
	var safeCount = 0

	// Process each set of data
	for _, data := range acc.rawData {
		var prevLevel = 0
		var levelDirection = 0
		var safe = true

		// Process each data item
	levels: // label to break out of loop early
		for _, level := range data {
			// handle first value
			if prevLevel == 0 {
				prevLevel = level
			} else {
				safe, levelDirection = checkSafety(level, prevLevel, levelDirection)
				prevLevel = level
			}
			if !safe {
				break levels
			}
		}
		if safe {
			safeCount++
		}
	}
	fmt.Println("Solution 1: " + strconv.Itoa(safeCount))
}

// Part 2 solution
func solution2(acc *SolutionAccumulator) {
}

// Checks the safety of the level data by 2 criteria
// 1. Are the levels steadily increasing or decreasing (can't change direction)
// 2. Is the change in level from the prior value between 1 and 3
func checkSafety(level int, prevLevel int, levelDirection int) (bool, int) {
	// If we haven't figured out the direction of the level, do that now.
	if levelDirection == 0 && prevLevel < level {
		levelDirection = 1
	} else if levelDirection == 0 && prevLevel > level {
		levelDirection = -1
	}

	var levelDiff = level - prevLevel
	var levelRange = absInt(levelDiff)

	// Check that the difference is within range
	if levelRange < 1 || levelRange > 3 {
		return false, levelDirection
	}

	// Confirm the direction is consistent
	return (levelDirection > 0 && levelDiff >= levelDirection) || (levelDirection < 0 && levelDiff <= levelDirection),
		levelDirection
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
