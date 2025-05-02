package LanguageDetection

import (
	"fmt"
)

func RunLanguage(filename string) {
	PathToFiles := "LanguageDetection/"
	content := ReadFile(PathToFiles + filename)
	fmt.Printf("The detected language of the file %v is %v\n", filename, LanguageDetection(content))
}

func Demo() {
	RunLanguage("English.txt")
	RunLanguage("French.txt")
	RunLanguage("German.txt")
}
