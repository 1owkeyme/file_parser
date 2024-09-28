package service

import (
	"file_parser/interfaces"
	"fmt"
	"math"
)

type StudentService struct {
	reader interfaces.Reader
	writer interfaces.Writer
}

func NewStudentService(r interfaces.Reader, w interfaces.Writer) *StudentService {
	return &StudentService{
		reader: r,
		writer: w,
	}
}

func (ss *StudentService) CalculateAverages(data map[string]map[string][]int) map[string]float64 {
	averages := make(map[string]float64)
	for name, subjects := range data {
		subjectCount := 0
		totalSum := 0.0
		for _, scores := range subjects {
			if len(scores) > 0 {
				average := ss.calculateAverage(scores)
				totalSum += average
				subjectCount++
			}
		}
		if subjectCount > 0 {
			averages[name] = totalSum / float64(subjectCount)
		}
	}
	return averages
}

func (ss *StudentService) calculateAverage(scores []int) float64 {
	sum := 0
	for _, score := range scores {
		sum += score
	}
	average := float64(sum) / float64(len(scores))
	return math.Round(average*100) / 100
}

func (ss *StudentService) Process() {
	data, err := ss.reader.Read()
	if err != nil {
		fmt.Println("error while reading data:", err)
		return
	}

	averages := ss.CalculateAverages(data)

	ss.writer.Write(data, averages)
}
