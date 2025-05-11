package LanguageDetection

import (
	"fmt"
)

// Programmierübung zu Clojure Woche 2, 7. Aufgabe

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
