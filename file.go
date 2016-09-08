package imago

import (
	"io/ioutil"
	"os"
)

// Read
func (f *ifile) Read(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

// Write
func (f *ifile) Write(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}
