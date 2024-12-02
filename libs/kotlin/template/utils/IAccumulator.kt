package utils

// Needs to be moved out to a kotlin-libs project
interface IAccumulator {
    fun processLine(line: String)
    fun execute()
}