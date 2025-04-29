package language_detection

import (
	"fmt"
)

func LanguageDetection(text string) string {
	// var de = []float32{6.51, 1.89, 3.06, 5.08, 17.40, 1.66, 3.01, 4.76, 7.55, 0.27, 1.21, 3.44, 2.53, 9.78, 2.51, 0.79, 0.02, 7.00, 7.27, 6.15, 4.35, 0.67, 1.89, 0.03, 0.04, 1.13}
	// var en = []float32{8.17, 1.48, 2.78, 4.25, 12.70, 2.23, 2.02, 6.09, 6.97, 0.15, 0.77, 4.03, 2.41, 6.75, 7.51, 1.93, 0.10, 5.99, 6.33, 9.06, 2.76, 0.98, 2.36, 0.15, 1.97, 0.07}
	// var fr = []float32{7.64, 0.90, 3.26, 3.67, 14.00, 1.07, 0.87, 0.74, 7.53, 0.55, 0.05, 5.46, 2.97, 7.10, 5.38, 3.02, 1.36, 6.55, 7.95, 7.24, 6.31, 1.63, 0.11, 0.39, 0.31, 0.14}
	fmt.Println("Hello world")

	var occurance = [26]int{}

	for _, char := range text {
		fmt.Println(char)
		if char > 64 && char < 91 {
			occurance[char-65] = occurance[char-65] + 1
		}
		if char > 96 && char < 123 {
			occurance[char-97] = occurance[char-97] + 1
		}
	}
	fmt.Println(occurance)

	return "worked"
}
