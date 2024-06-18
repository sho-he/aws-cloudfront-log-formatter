package fileoperators

import "os"

type FileOperator struct {
	reader  Reader
	creater Creater
}

func NewFileOperator(reader Reader, creater Creater) *FileOperator {
	return &FileOperator{
		reader,
		creater,
	}
}

func (f *FileOperator) ReadFile() (*[]byte, error) {
	return f.reader.ReadFile()
}

func (f *FileOperator) CreateFile() (*os.File, error) {
	return f.creater.CreateFile()
}
