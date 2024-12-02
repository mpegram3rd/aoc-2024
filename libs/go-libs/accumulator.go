package utils

// Interface for a line by line accumulator
type IAccumulator interface {
	ProcessLine(string)
	Execute()
}
