package imago

import (
	"io/ioutil"
	"os"
)

// file file
type file struct{}

// File new file
var File = new(file)

// Read file read
// date 2016-12-31
// author andy.jiang
func (f file) Read(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

// Write file write
// date 2016-12-31
// author andy.jiang
func (f file) Write(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}
