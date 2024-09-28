package writer

import (
	"file_parser/interfaces"
	"fmt"
	"os"
	"sort"
	"strings"
)

type ConsoleWriter struct{}

var _ interfaces.Writer = (*ConsoleWriter)(nil)

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{}
}

func (cw *ConsoleWriter) Write(data map[string]map[string][]int, averages map[string]float64) {
	output := os.Stdout

	names := make([]string, 0, len(data))
	for name := range data {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		if averages[names[i]] != averages[names[j]] {
			return averages[names[i]] > averages[names[j]]
		}
		return names[i] < names[j]
	})

	for _, name := range names {
		scores := data[name]
		fmt.Fprintln(output, name)
		for subject, scores := range scores {
			scoresStr := strings.Join(intSliceToStringSlice(scores), ", ")
			fmt.Fprintf(output, "%s scores: %s\n", subject, scoresStr)
		}
		fmt.Fprintf(output, "Average score: %.2f\n\n", averages[name])
	}
}
