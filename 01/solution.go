package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Struct to accumulate data from the input file
type SolutionAccumulator struct {
	list1 []int
	list2 []int
}

// Process and accumulate data from the file
func (acc *SolutionAccumulator) ProcessLine(line string) {
	parts := strings.Split(line, "   ")
	val1, _ := strconv.Atoi(parts[0])
	val2, _ := strconv.Atoi(parts[1])
	acc.list1 = append(acc.list1, val1)
	acc.list2 = append(acc.list2, val2)
}

// Execute the strategies for each solution
func (acc *SolutionAccumulator) Execute() {
	sort.Slice(acc.list1, func(i, j int) bool { return acc.list1[i] < acc.list1[j] })
	sort.Slice(acc.list2, func(i, j int) bool { return acc.list2[i] < acc.list2[j] })
	solution1(acc)
	solution2(acc)
}

// Part 1 solution
func solution1(acc *SolutionAccumulator) {
	var total = 0.0
	for idx, element := range acc.list1 {
		total += math.Abs(float64(element) - float64(acc.list2[idx]))
	}
	fmt.Println("Solution 1: " + strconv.FormatFloat(total, 'f', 1, 64))
}

// Part 2 solution
// This does some work to keep track of where the last match was found in the right hand list
// since the lists are both sorted there's no reason to go back to the beginning to find matches
func solution2(acc *SolutionAccumulator) {
	var length = len(acc.list1) - 1
	var total = 0.0
	var idx2 = 0
	var mostRecentVal = 0.0
	var mostRecentRepeatsLeft = 0.0
	var repeatsRight = 0.0

	// Loop through all the ones in the list
	for idx, element := range acc.list1 {
		var currentVal = float64(element)
		if currentVal > mostRecentVal || idx == length {
			total += mostRecentVal * mostRecentRepeatsLeft * repeatsRight
			mostRecentVal = currentVal
			mostRecentRepeatsLeft = 1
			repeatsRight, idx2 = findRepeats(mostRecentVal, acc.list2, idx2)
		} else {
			mostRecentRepeatsLeft += 1.0
		}
	}
	// Add in the final set if there are others left over.
	if repeatsRight > 0 {
		total += mostRecentVal * mostRecentRepeatsLeft * repeatsRight
	}
	fmt.Println("Solution 2: " + strconv.FormatFloat(total, 'f', 1, 64))
}

// Finds the number of repeating numbers in the right hand list.  Always starts from the last matching
// index value of the prior matches so we don't have to re-process the whole right hand list.
func findRepeats(val float64, list2 []int, idx int) (float64, int) {
	var startIdx = idx
	var done = false
	var maxLen = len(list2)
	var matchCount = 0.0

	// Keep processing until we reach the end of the list or we find a value higher than the one we want to match
	for startIdx < maxLen && !done {
		var currentVal = float64(list2[startIdx])
		if currentVal == val {
			matchCount += 1.0
		}
		startIdx++
		done = currentVal > val
	}
	// If we hit the end of the list that means we didn't really find a value
	// So reset the start index back to its original position
	if done && matchCount == 0.0 {
		startIdx = idx
	} else {
		startIdx -= 1
	}

	return matchCount, startIdx
}
