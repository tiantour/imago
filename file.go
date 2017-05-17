package imago

import (
	"io/ioutil"
	"os"
)

// File file
var File = &iFile{}

type iFile struct{}

// Read
func (f *iFile) Read(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

// Write
func (f *iFile) Write(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}
