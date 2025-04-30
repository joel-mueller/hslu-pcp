package LanguageDetection

import (
	"math"
)

func cosineSimilarity(list1 [26]float64, list2 [26]float64) float64 {
	var up float64 = 0.0
	for i := 0; i < len(list1); i++ {
		up += list1[i] * list2[i]
	}
	return up / (pythagoreanList(list1) * pythagoreanList(list2))
}

func pythagoreanList(list [26]float64) float64 {
	var result float64 = 0.0
	for _, i := range list {
		result += i * i
	}
	return math.Sqrt(result)
}

func LanguageDetection(text string) string {
	var de = [26]float64{6.51, 1.89, 3.06, 5.08, 17.40, 1.66, 3.01, 4.76, 7.55, 0.27, 1.21, 3.44, 2.53, 9.78, 2.51, 0.79, 0.02, 7.00, 7.27, 6.15, 4.35, 0.67, 1.89, 0.03, 0.04, 1.13}
	var en = [26]float64{8.17, 1.48, 2.78, 4.25, 12.70, 2.23, 2.02, 6.09, 6.97, 0.15, 0.77, 4.03, 2.41, 6.75, 7.51, 1.93, 0.10, 5.99, 6.33, 9.06, 2.76, 0.98, 2.36, 0.15, 1.97, 0.07}
	var fr = [26]float64{7.64, 0.90, 3.26, 3.67, 14.00, 1.07, 0.87, 0.74, 7.53, 0.55, 0.05, 5.46, 2.97, 7.10, 5.38, 3.02, 1.36, 6.55, 7.95, 7.24, 6.31, 1.63, 0.11, 0.39, 0.31, 0.14}

	occurrence := getOccurrence(text)

	var a = map[string]float64{"de": cosineSimilarity(occurrence, de), "en": cosineSimilarity(occurrence, en), "fr": cosineSimilarity(occurrence, fr)}

	var max_key = "no key here"
	var max_value = 0.0
	for key, value := range a {
		if value > max_value {
			max_key = key
			max_value = value
		}
	}
	return max_key
}

func getOccurrence(text string) [26]float64 {
	var occurrence = [26]float64{}

	for _, char := range text {
		if char > 64 && char < 91 {
			occurrence[char-65] = occurrence[char-65] + 1
		}
		if char > 96 && char < 123 {
			occurrence[char-97] = occurrence[char-97] + 1
		}
	}
	return occurrence
}
