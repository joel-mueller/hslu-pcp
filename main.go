package main

import (
	"fmt"
	"hslu-pcp/advent_of_code"
	"hslu-pcp/language_detection"
	//"hslu-pcp/language_detection"
	//"hslu-pcp/advent_of_code"
)

func run_advent_of_code(steps int) int {
	start := []int{3, 4, 3, 1, 2}
	return advent_of_code.Advent(start, steps)
}

func run_language_detection(text string) string {
	return language_detection.LanguageDetection(text)
}

func main() {

	fmt.Println(run_advent_of_code(50))
	fmt.Println(run_language_detection("The origins of the name Java are not clear. The island could possibly have been named after the jáwa-wut plant, which was said to be common in the island during the time, and that prior to Indianization the island had different names.[5] There are other possible sources: the word jaú and its variations mean beyond or distant.[6] And, in Sanskrit yava means barley, a plant for which the island was famous.[6] Yavadvipa is mentioned in India's earliest epic, the Ramayana. Sugriva, the chief of Rama's army, dispatched his men to Yavadvipa, the island of Java, in search of Sita.[7] It was hence referred to in India by the Sanskrit name yāvaka dvīpa (dvīpa = island). Java is mentioned in the ancient Tamil text Manimekalai by Chithalai Chathanar which states that Java had a kingdom with a capital called Nagapuram.[8][9][10] Another source states that the word Java is derived from a Proto-Austronesian root word, meaning home.[11] The great island of Iabadiu or Jabadiu was mentioned in Ptolemy's Geographia composed around 150 CE in the Roman Empire. Iabadiu is said to mean barley island, to be rich in gold, and have a silver town called Argyra at the west end. The name indicates Java[12] and seems to be derived from the Sanskrit name Java-dvipa (Yavadvipa)."))
}
