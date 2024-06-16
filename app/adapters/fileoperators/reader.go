package fileoperators

type Reader interface {
	ReadFile() (*[]byte, error)
}
