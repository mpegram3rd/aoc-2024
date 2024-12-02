package utils

import (
	"bufio"
	"log"
	"os"
)

// Interface for file processor so it can be passed around
type IFileProcessor interface {
	processFile()
}

// File processor structure that we can hang an IFileProcessor implementation off of
type FileProcessor struct {
	Filename    string
	Accumulator IAccumulator
}

// Opens and iterates through a file using the attached IAccumulator
func (fp FileProcessor) ProcessFile() {
	// open file and check for errors
	file, err := os.Open(fp.Filename)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	garbIn := bufio.NewScanner(file)
	for garbIn.Scan() {
		fp.Accumulator.ProcessLine(garbIn.Text())
	}

	// Run any final processing after the file has been processed.
	fp.Accumulator.Execute()
}
