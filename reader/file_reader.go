package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileReader struct {
	filePath string
}

func NewFileReader(filepath string) *FileReader {
	return &FileReader{filePath: filepath}
}

func (fr *FileReader) Read() (map[string][]int, error) {
	f, err := os.Open(fr.filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	students := make(map[string][]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}
		name := parts[0]
		score, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid score for student %s: %s", name, parts[1])
		}
		students[name] = append(students[name], score)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return students, nil
}
