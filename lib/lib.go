package lib

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// FileExists checks if the file exists
func FileExists(fileName string) (bool, error) {
	filepath, _ := filepath.Abs(fileName)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false, err
	}

	return true, nil
}

// LoadFile loads a yaml file into a byte slice
func LoadFile(fileName string) []byte {
	filepath, _ := filepath.Abs(fileName)
	yamlFile, ymlErr := ioutil.ReadFile(filepath)
	if ymlErr != nil {
		log.Println(ymlErr)
		log.Println(yamlFile)
	}

	return yamlFile
}
