package ports

import "os"

type FileOperator interface {
	ReadFile() (*[]byte, error)
	CreateFile() (*os.File, error)
}
