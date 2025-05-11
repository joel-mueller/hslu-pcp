package DELETEME

import (
	"fmt"
)

func Demooo() {
	safeRun(outOfBounds)
	fmt.Println("Program continues after panic is recovered.")
	safeRun(myPanic)
	fmt.Println("Program continues after panic is recovered.")
}

func safeRun(method func()) {
	defer handlePanic()
	method()
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func outOfBounds() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[5])
}

func myPanic() {
	panic("something bad happened")
}
