package main

import (
	"bufio"
	"fmt"
	"hslu-pcp/AdventOfCode"
	"hslu-pcp/Datastructures"
	"hslu-pcp/FunctionalPatterns"
	"hslu-pcp/LanguageDetection"
	"hslu-pcp/Panic"
	"hslu-pcp/Routines"
	"hslu-pcp/RoutinesChannels"
	"os"
	"strings"
)

var demos = map[string]func(){
	"language":   LanguageDetection.Demo,
	"advent":     AdventOfCode.Demo,
	"stacks":     Datastructures.Demo,
	"functional": FunctionalPatterns.Demo,
	"routines":   Routines.Demo,
	"panic":      Panic.Demo,
	"channels":   RoutinesChannels.Demo,
}

func help() {
	fmt.Println("Available demos:")
	for name := range demos {
		fmt.Printf("- %s\n", name)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello! What program do you want to run?")
	help()

	for {
		fmt.Print("Enter demo name (or 'help' or 'exit'): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "exit":
			fmt.Println("Goodbye!")
			return
		case "help":
			help()
		default:
			if demo, ok := demos[input]; ok {
				fmt.Printf("Running demo: %s\n\n", input)
				demo()
				fmt.Println("\nDone.")
			} else {
				fmt.Println("Unknown demo. Type 'help' to see available demos.")
			}
		}
	}
}
