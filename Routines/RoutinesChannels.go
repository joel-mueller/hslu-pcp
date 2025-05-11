package Routines

import (
	"fmt"
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
				DoBlockingWait(500)
			}
		}
	}()

	fmt.Println(LastTask())

	close(stop)

	fmt.Print("-> Done.")
}

func DoBlockingWait(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

func LongLastingTask(c chan int) {
	DoBlockingWait(3000)
	fmt.Print("3000")
	c <- 3000
}

func EvenLongerLastingTask(c chan int) {
	DoBlockingWait(6000)
	fmt.Print("6000")
	c <- 6000
}

func LastTask() string {
	c := make(chan int)

	go LongLastingTask(c)
	go EvenLongerLastingTask(c)

	s, t := <-c, <-c

	DoBlockingWait(2000)
	return fmt.Sprintf("Was waiting for %v ms", s+t+2000)
}
