package main

import (
	"file_parser/interfaces"
	reader_impl "file_parser/reader"
	"file_parser/service"
	writer_impl "file_parser/writer"
	"flag"
)

func main() {
	var filePath, outputFile string
	flag.StringVar(&filePath, "f", "", "input file path")
	flag.StringVar(&outputFile, "o", "", "output file path")
	flag.Parse()

	var reader interfaces.Reader
	if filePath != "" {
		reader = reader_impl.NewFileReader(filePath)
	} else {
		reader = reader_impl.NewConsoleReader()
	}

	var writer interfaces.Writer
	if outputFile != "" {
		writer = writer_impl.NewFileWriter(outputFile)
	} else {
		writer = writer_impl.NewConsoleWriter()
	}
	service.NewStudentService(reader, writer).Process()
}
