package config

import (
	"fmt"
	"io/ioutil"
)

// Read a file from disk.
func read(file string) []byte {

	content, err := ioutil.ReadFile(configDir + file)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return content
}
