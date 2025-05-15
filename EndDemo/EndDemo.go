package EndDemo

import (
	"fmt"
	"sync"
	"time"
)

func Demo() {
	defer fmt.Println("Final Task finished.")

	var wg sync.WaitGroup
	wg.Add(3) // ← make room for three Done() calls

	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})

	go func() {
		defer wg.Done()
		TaskOne()
		close(c1)
	}()
	go func() {
		defer wg.Done()
		TaskTwo()
		close(c2)
	}()
	go func() {
		defer wg.Done()
		TaskThree() // panic in here will be recovered
		close(c3)
	}()

	wg.Wait() // ← block until all three Done()s have run
}

func TaskOne() {
	time.Sleep(2 * time.Second)
	fmt.Println("Task 1 is done.")
}

func TaskTwo() {
	time.Sleep(4 * time.Second)
	fmt.Println("Task 2 is done.")
}

func TaskThree() {
	time.Sleep(4 * time.Second)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Task 3:", r)
		}
	}()
	panic("something went wrong in Task 3")
}
