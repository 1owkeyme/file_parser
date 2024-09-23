package writer

import (
	"fmt"
	"sort"
	"strings"
)

type ConsoleWriter struct{}

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{}
}

func (cw *ConsoleWriter) Write(data map[string][]int, averages map[string]float64) {
	names := make([]string, 0, len(data))
	for name := range data {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		scores := data[name]
		fmt.Println(name)
		fmt.Printf("Scores: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(scores)), ", "), "[]"))
		fmt.Printf("Average score: %.2f\n", averages[name])
		fmt.Println()
	}
}
