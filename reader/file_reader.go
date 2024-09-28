package reader

import (
	"bufio"
	"file_parser/interfaces"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileReader struct {
	filePath string
}

var _ interfaces.Reader = (*FileReader)(nil)

func NewFileReader(filepath string) *FileReader {
	return &FileReader{filePath: filepath}
}

func (fr *FileReader) Read() (map[string]map[string][]int, error) {
	f, err := os.Open(fr.filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	students := make(map[string]map[string][]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		name, subject := parts[0], parts[1]
		if _, valid := ValidSubjects[subject]; !valid {
			fmt.Printf("Warning: invalid subject for student %s: %s\n", name, subject)
			continue
		}

		score, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("invalid score for student %s: %s", name, parts[1])
		}
		if students[name] == nil {
			students[name] = make(map[string][]int)
		}

		students[name][subject] = append(students[name][subject], score)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return students, nil
}
