package interfaces

type Writer interface {
	Write(data map[string]map[string][]int, averages map[string]float64)
}
