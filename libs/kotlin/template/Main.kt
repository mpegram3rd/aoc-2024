import utils.FileProcessor

fun main(args: Array<String>) {
    val fileProcessor = FileProcessor(args[0], Accumulator())
    fileProcessor.processFile()
}