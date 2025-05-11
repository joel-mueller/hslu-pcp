package DELETEME

import (
	"fmt"
	"sync"
	"time"
)

// PCP Aufgabe 3 von SW09

func Demo() {
	fmt.Println("->Now waiting for things to happen")

	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				fmt.Print(".")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	RunCombinedTasks(stop)
}

func LongLastingTask() int64 {
	time.Sleep(3 * time.Second)
	fmt.Print("3000")
	return 3000
}

func EvenLongerLastingTask() int64 {
	time.Sleep(6 * time.Second)
	fmt.Print("6000")
	return 6000
}

func LastTask() {
	time.Sleep(2 * time.Second)
}

func RunCombinedTasks(stop chan struct{}) {
	var wg sync.WaitGroup
	var result1, result2 int64

	wg.Add(2)

	go func() {
		defer wg.Done()
		result1 = LongLastingTask()
	}()

	go func() {
		defer wg.Done()
		result2 = EvenLongerLastingTask()
	}()

	wg.Wait()

	LastTask()

	fmt.Printf("was waiting for %dms", result1+result2+2000)
	fmt.Println()

	close(stop)
	fmt.Print("-> Done.")
}
