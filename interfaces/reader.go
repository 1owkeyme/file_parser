package interfaces

type Reader interface {
	Read() (map[string]map[string][]int, error)
}
