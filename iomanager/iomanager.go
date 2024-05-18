package iomanager

type IOManager interface {
	ReadFile() ([]string, error)
	WriteResult(content interface{}) error
}