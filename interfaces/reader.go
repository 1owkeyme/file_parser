package interfaces

type Reader interface {
	Read() (map[string][]int, error)
}
