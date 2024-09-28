package writer

import (
	"file_parser/interfaces"
	"fmt"
	"os"
	"sort"
	"strings"
)

type FileWriter struct {
	outputFile string
}

var _ interfaces.Writer = (*ConsoleWriter)(nil)

func NewFileWriter(outputFile string) *FileWriter {
	return &FileWriter{outputFile: outputFile}
}

func (fw *FileWriter) Write(data map[string]map[string][]int, averages map[string]float64) {
	output, err := os.Create(fw.outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer output.Close()

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
