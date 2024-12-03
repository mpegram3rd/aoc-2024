package main

import (
	"blt.com/utils"
	"os"
)

func main() {

	filePath := os.Args[1] // read file path from commandline

	solutionFP := utils.FileProcessor{Filename: filePath, Accumulator: &SolutionAccumulator{}}
	solutionFP.ProcessFile()

}
