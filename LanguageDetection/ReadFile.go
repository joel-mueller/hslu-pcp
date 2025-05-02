package LanguageDetection

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error Occurred while trying to clone the file", err)
		}
	}(file)

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	return string(bytes)
}
