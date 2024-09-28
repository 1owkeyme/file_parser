package reader

import (
	"bufio"
	"file_parser/interfaces"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ConsoleReader struct{}

var _ interfaces.Reader = (*ConsoleReader)(nil)

func NewConsoleReader() *ConsoleReader {
	return &ConsoleReader{}
}

func (cr *ConsoleReader) Read() (map[string]map[string][]int, error) {
	students := make(map[string]map[string][]int)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter student records (format: Name Subject Score) or press Enter on an empty line to finish:")

	for {
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}

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
			return nil, fmt.Errorf("invalid score for student %s: %s", name, parts[2])
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
