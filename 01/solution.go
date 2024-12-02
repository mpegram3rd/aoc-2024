package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type SolutionAccumulator struct {
	list1 []int
	list2 []int
}

func (acc *SolutionAccumulator) ProcessLine(line string) {
	parts := strings.Split(line, "   ")
	val1, _ := strconv.Atoi(parts[0])
	val2, _ := strconv.Atoi(parts[1])
	acc.list1 = append(acc.list1, val1)
	acc.list2 = append(acc.list2, val2)
}

func (acc *SolutionAccumulator) Execute() {
	sort.Slice(acc.list1, func(i, j int) bool { return acc.list1[i] < acc.list1[j] })
	sort.Slice(acc.list2, func(i, j int) bool { return acc.list2[i] < acc.list2[j] })
	solution1(acc)
}

func solution1(acc *SolutionAccumulator) {
	var total = 0.0
	for idx, element := range acc.list1 {
		total += math.Abs(float64(element) - float64(acc.list2[idx]))
	}
	fmt.Println("Solution 1: " + strconv.FormatFloat(total, 'f', 1, 64))
}
