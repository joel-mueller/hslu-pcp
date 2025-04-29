package main

import (
	"fmt"
	"hslu-pcp/language_detection"
)

func main() {
	fmt.Println(string(rune(65)))
	fmt.Println(language_detection.LanguageDetection("aAZz"))
}
