package main

import (
	"file_parser/reader"
	"file_parser/service"
	"file_parser/writer"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("provide path to file as first argument")
		return
	}

	filePath := os.Args[1]

	fileReader := reader.NewFileReader(filePath)
	consoleWriter := writer.NewConsoleWriter()
	service.NewStudentService(fileReader, consoleWriter).Process()
}
