package Panic

import (
	"log"
)

func Demo() {
	safeRun(outOfBounds)
	log.Println("Program continues after panic is recovered.")
	safeRun(myPanic)
	log.Println("Program continues after panic is recovered.")
}

func safeRun(method func()) {
	defer handlePanic()
	method()
}

func handlePanic() {
	if r := recover(); r != nil {
		log.Println("Recovered from panic:", r)
	}
}

func outOfBounds() {
	arr := []int{1, 2, 3}
	log.Println(arr[5])
}

func myPanic() {
	panic("something bad happened")
}
