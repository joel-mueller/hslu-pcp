package Weather

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// Programmierübung zu Modern Java Woche 2, 2. Aufgabe

func Demo() {
	numServices := 3
	weather := make(chan string, numServices)
	failures := make(chan struct{}, numServices)
	for i := 1; i <= numServices; i++ {
		go func(n int) {
			defer func() {
				if r := recover(); r != nil {
					//fmt.Println("Recovered Task", r)
					failures <- struct{}{}
				}
			}()
			CallWeatherService(n, weather)
		}(i)
	}
	failed := 0
	for {
		select {
		case result := <-weather:
			fmt.Println("Weather empfangen:", result)
			return
		case <-failures:
			failed++
			if failed == numServices {
				fmt.Println("Alle Wetterdienste sind fehlgeschlagen.")
				return
			}
		}
	}
}

func CallWeatherService(number int, c chan string) {
	delay := 200 + rand.IntN(800)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	if rand.IntN(2) == 1 {
		panic(fmt.Sprintf("Wetter Service %v ist felgeschlagen", number))
	}
	c <- fmt.Sprintf("Wetter Service %v war erfolgreich", number)
}
