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
	solution1(acc)
	solution2(acc)
}

// Part 1 solution
func solution1(acc *SolutionAccumulator) {
	var safeCount = 0

	// Process each set of data
	for ri, data := range acc.rawData {
		var prevLevel = 0
		//		var levelDirection = 0
		//		var safe = true

		// Process each data item
		var positives = 0
		var negatives = 0
		var outOfRange = 0
		for _, level := range data {
			// handle first value
			if prevLevel == 0 {
				prevLevel = level
			} else {
				if prevLevel < level {
					positives++
				} else if level < prevLevel {
					negatives++
				}
				var diff = absInt(level - prevLevel)
				if diff < 1 || diff > 3 {
					outOfRange++
				}
				prevLevel = level
			}
		}
		if checkSafety(positives, negatives, outOfRange, 0) {
			safeCount++
			reportSafe(ri, data)
		}
	}
	fmt.Println("Solution 1: " + strconv.Itoa(safeCount))
}

// Part 2 solution
func solution2(acc *SolutionAccumulator) {
}

func reportSafe(ri int, data []int) {
	fmt.Printf("Row %d: ", ri)
	for _, val := range data {
		fmt.Printf("%d, ", val)
	}
	fmt.Println()
}
func checkSafety(positives int, negatives int, outOfRange int, allowedFailures int) bool {
	var failureCount = 0

	if positives > negatives {
		failureCount += negatives // the lesser value indicates the number of directional switches
	} else {
		if positives < negatives {
			failureCount += positives
		} else {
			failureCount += positives // if they're equal then just pick one
		}
	}

	failureCount += outOfRange
	return failureCount <= allowedFailures
}

// Checks the safety of the level data by 2 criteria
// 1. Are the levels steadily increasing or decreasing (can't change direction)
// 2. Is the change in level from the prior value between 1 and 3
//func checkSafety(level int, prevLevel int, levelDirection int) (bool, int) {
//	// If we haven't figured out the direction of the level, do that now.
//	if levelDirection == 0 && prevLevel < level {
//		levelDirection = 1
//	} else if levelDirection == 0 && prevLevel > level {
//		levelDirection = -1
//	}
//
//	var levelDiff = level - prevLevel
//	var levelRange = absInt(levelDiff)
//
//	// Check that the difference is within range
//	if levelRange < 1 || levelRange > 3 {
//		return false, levelDirection
//	}
//
//	// Confirm the direction is consistent
//	return (levelDirection > 0 && levelDiff >= levelDirection) || (levelDirection < 0 && levelDiff <= levelDirection),
//		levelDirection
//}

// Determines which direction the levels are trending: -1 = down, 0 = has not been determined, 1 = up
//func determineDirection(prevLevel int, level int) {
//	var levelDirection = 0
//	// If we haven't figured out the direction of the level, do that now.
//	if prevLevel < level {
//		levelDirection = 1
//	} else if prevLevel > level {
//		levelDirection = -1
//	}
//
//}

// An integer based Absolute Value function
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
