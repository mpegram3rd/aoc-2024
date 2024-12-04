package main

import (
	"fmt"
	//	"math"
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
	fmt.Println("Starting processing!")
	// checkLevels(acc, 0) // Solution 1
	checkLevels(acc, 1) // Solution 2
}

// Part 1 solution
func checkLevels(acc *SolutionAccumulator, maxViolations int) {
	var safeCount = 0

	// Process each set of data
	for ri, data := range acc.rawData {
		var prevLevel = 0
		var violationCount = 0
		var direction = findGeneralDirection(data)
		for _, level := range data {
			if prevLevel != 0 {
				var diff = level - prevLevel
				var absDiff = absInt(diff)
				if direction == 1 && diff < 1 { // if ascending and we hit a non-ascending value
					fmt.Println("Failure: Ascending with change of direction")
					violationCount++
					level = prevLevel
				} else if direction == -1 && diff > -1 { // if descending and we hit an ascending value
					fmt.Println("Failure: Descending with change of direction")
					violationCount++
				} else if absDiff < 1 || absDiff > 3 {
					fmt.Println("Failure: Difference too large")
					violationCount++
					//level = prevLevel
				}
			}
			prevLevel = level

		}

		if violationCount <= maxViolations {
			safeCount++
			reportSafe(ri, data)
		}
	}
	fmt.Println("Solution: " + strconv.Itoa(safeCount))
}

// Part 2 solution
func reportSafe(ri int, data []int) {
	fmt.Printf("Row %d: ", ri)
	for _, val := range data {
		fmt.Printf("%d, ", val)
	}
	fmt.Println()
}

// Figure out which direction in general the list is headed.
func findGeneralDirection(data []int) int {
	var positives = 0
	var negatives = 0
	var prevVal = 0

	for _, val := range data {
		if prevVal != 0 {
			if val > prevVal {
				positives++
			} else if val < prevVal {
				negatives++
			}
		}
		prevVal = val
	}

	if positives > negatives {
		return 1
	}
	return -1
}

// Determines which direction the levels are trending: -1 = down, 0 = has not been determined, 1 = up
func determineDirection(prevLevel int, level int) int {
	var levelDirection = 0
	// If we haven't figured out the direction of the level, do that now.
	if prevLevel < level {
		levelDirection = 1
	} else if prevLevel > level {
		levelDirection = -1
	}
	return levelDirection
}

// An integer based Absolute Value function
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
