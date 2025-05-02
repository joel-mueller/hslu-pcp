package LanguageDetection

import (
	"math"
	"testing"
)

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

func TestPythagoreanList(t *testing.T) {
	list := [26]float64{3, 4}
	expected := 5.0 // sqrt(3^2 + 4^2)
	result := pythagoreanList(list)
	if !floatEquals(result, expected) {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCosineSimilarity_Identical(t *testing.T) {
	list := [26]float64{1, 2, 3, 4, 5}
	result := cosineSimilarity(list, list)
	if !floatEquals(result, 1.0) {
		t.Errorf("Expected cosine similarity of 1.0, got %f", result)
	}
}

func TestCosineSimilarity_Orthogonal(t *testing.T) {
	list1 := [26]float64{1, 0}
	list2 := [26]float64{0, 1}
	result := cosineSimilarity(list1, list2)
	if !floatEquals(result, 0.0) {
		t.Errorf("Expected cosine similarity of 0.0, got %f", result)
	}
}

func TestGetOccurrence_Lowercase(t *testing.T) {
	text := "abc"
	result := getOccurrence(text)
	expected := [26]float64{1, 1, 1}
	for i := 0; i < 3; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %f at index %d, got %f", expected[i], i, result[i])
		}
	}
}

func TestGetOccurrence_Uppercase(t *testing.T) {
	text := "ABC"
	result := getOccurrence(text)
	expected := [26]float64{1, 1, 1}
	for i := 0; i < 3; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %f at index %d, got %f", expected[i], i, result[i])
		}
	}
}

func TestLanguageDetection_English(t *testing.T) {
	text := "This is a simple English sentence.Java[a] is one of the Greater Sunda Islands in Indonesia. It is bordered by the Indian Ocean to the south and the Java Sea (a part of Pacific Ocean) to the north. With a population of 156.9 million people (including Madura) in mid 2024, projected to rise to 158 million at mid 2025, Java is the world's most populous island, home to approximately 55.7% of the Indonesian population (only approximately 44.3% of Indonesian population live outside Java).[2] Indonesia's capital city, Jakarta, is on Java's northwestern coast."
	lang := LanguageDetection(text)
	if lang != "en" {
		t.Errorf("Expected 'en', got %s", lang)
	}
}

func TestLanguageDetection_German(t *testing.T) {
	text := "Dies ist ein einfacher deutscher Satz.Java, indonesisch Jawa (nach alter Schreibweise Djawa; Aussprache: [dʒawa], im Deutschen zumeist [ˈjaːva]) ist neben Sumatra, Borneo und Sulawesi eine der vier Großen Sundainseln. Die Insel gehört vollständig zur Republik Indonesien, auf ihr liegt auch die größte Stadt und ehemalige Hauptstadt Indonesiens, Jakarta.Weiter unten sind drei Häufigkeitsvektoren für die Sprachen Deutsch, Englisch und Französisch gegeben. Schreiben Sie eine Funktion, die mit den Funktionen aus a.) und b.) für ein Stück Text berechnet, welcher der drei Häufigkeitsvektoren am ähnlichsten ist und dann die wahrscheinlichste Sprache ausgibt (z.B. «de», «en» oder «fr»)."
	lang := LanguageDetection(text)
	if lang != "de" {
		t.Errorf("Expected 'de', got %s", lang)
	}
}

func TestLanguageDetection_French(t *testing.T) {
	text := "Ceci est une phrase française simple.Java, en indonésien Jawa, Djawa jusqu'à la réforme orthographique de 1972, est une île du Sud-Ouest de l'Indonésie faisant partie de l'Insulinde. Elle se situe en mer de Java, qui la baigne à l'ouest (détroit de la Sonde) et au nord. Elle est aussi bordée à l'est par la mer de Bali et au sud par l'océan Indien. Elle est entourée par les îles de Sumatra au nord-ouest, Bali à l'est et Bornéo au nord. Son nom viendrait du sanskrit Javadvipa, « l'île du millet ». C'est sous ce nom que l'île est en effet désignée dans l'épopée indienne du Ramayana (écrite entre le"
	lang := LanguageDetection(text)
	if lang != "fr" {
		t.Errorf("Expected 'fr', got %s", lang)
	}
}

func TestLanguageDetection_Empty(t *testing.T) {
	text := ""
	lang := LanguageDetection(text)
	if lang != "no key here" {
		t.Errorf("Expected 'no key here', got %s", lang)
	}
}
