package imago

import (
	"io/ioutil"
	"os"
)

// File file
type File struct{}

// NewFile new file
func NewFile() *File {
	return &File{}
}

// Read file read
// date 2016-12-31
// author andy.jiang
func (f File) Read(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

// Write file write
func (f File) Write(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}
