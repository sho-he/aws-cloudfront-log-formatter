package fileoperators

import "os"

type Creater interface {
	CreateFile() (*os.File, error)
}
