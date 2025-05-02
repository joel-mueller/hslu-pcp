package LanguageDetection

import "fmt"

func Demo() {
	PathToFiles := "LanguageDetection/"
	content := ReadFile(PathToFiles + "English.txt")
	//fmt.Println(content)
	fmt.Println(LanguageDetection(content))
}
