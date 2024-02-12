package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(any) error
}
