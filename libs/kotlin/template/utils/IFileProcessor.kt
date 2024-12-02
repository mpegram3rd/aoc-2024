package utils

// Needs to be moved out to a kotlin-libs project
interface IFileProcessor {
    val filename : String
    val accumulator : IAccumulator
    fun processFile()
}