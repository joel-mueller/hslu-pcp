package EndDemo

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Demo() {
	defer fmt.Println("Final Task finished.")

	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup
	wg.Add(3)

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
		TaskThree()
		close(c3)
	}()

	wg.Wait()
}

func TaskOne() {
	time.Sleep(2 * time.Second)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Task 1:", r)
		}
	}()
	if rand.Intn(2) == 1 {
		panic("Task 1 failed unexpectedly")
	}
	fmt.Println("Task 1 is done.")
}

func TaskTwo() {
	time.Sleep(4 * time.Second)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Task 2:", r)
		}
	}()
	if rand.Intn(2) == 1 {
		panic("Task 2 failed unexpectedly")
	}
	fmt.Println("Task 2 is done.")
}

func TaskThree() {
	time.Sleep(4 * time.Second)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Task 3:", r)
		}
	}()
	if rand.Intn(2) == 1 {
		panic("Task 3 failed unexpectedly")
	}
	fmt.Println("Task 3 is done.")
}
