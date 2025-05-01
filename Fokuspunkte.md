# Fokuspunkte

- Defer, panic und recover
- Goroutines, Channels & Select
- Structural & Nominal Typing
- The Go Memory Model

Abgabe: Ihre gewählten 3 bis 7 Fokuspunkte, jeweils
inkl. kurze Erklärung (stichwortartig, ggf. inkl.kurze
Pseudo-Code-Sequenzen)

## `The Go Memory Model`

### `Goroutines and Concurrency`

- Go verwendet *goroutines* für leichtgewichtige parallelität.
- Speicher Zugriff zwischen *goroutines* muss synchronisiert werden um race conditions zu veremeiden.

### `Happens-before Relationship`

- Das Memory Model definiert welche operationen garantiert sichtbar für andere *goroutinen* ist.
- Falls eine Aktion vor einer anderen passiert, dann wird garantiert das der nächste diese sieht.

### `Atomic Operations`
- Das ```sync/atomic``` package bietet low-level atomaren memory Zugriff mit garantierter visibility.

### `Compiler and CPU Reordering`
- Das Model erlaubt Compilers und CPUs Instruktionen um zu ordnen solange das happens-before Regel respektiert.

### `Data Races are Bugs`
- Go behandelt Data races als programmier Errors.



## `goroutines, channels & select`
### `goroutine`
Eine goroutine beschreibt einen leichtgewichtigen thread welcher von der Go Runtime gemanaged wird.
Beispiel:
```go
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
```



### `channels`
Channels sind eine typisierte Leitung durch welche man Daten senden und bekommen kann durch den channel operator '<-'.

Wie maps und slices müssen channels kreiirt werden bevor diese verwendet werden können

```go
ch := make(chan int)
```

Beispiel:

```go
ch <- v // Send v to channel ch
v := <-ch // Receive from ch, and assign value to v
```

```go
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```

### `select`
Das select statement lässt eine goroutine auf mehrere kommunikations operationen warten.

Ein select blockiert solange bis einer seiner cases ausführbar ist, dann führt es diesen aus. Falls mehrere gleichzeitig ausführbar sind wird zufällig ausgewählt.

```go
func main() {
ch1 := make(chan string)
ch2 := make(chan string)

    // Simulate concurrent operations
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Message from channel 1"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Message from channel 2"
    }()

    // Use select to wait for either channel
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received:", msg2)
        }
    }
}
```

## `defer`, `panic` und `recover`

### Wann braucht man `panic` ?

`panic` wird gebraucht wenn etwas grosses unerwartetes vorliegt und das Programm nicht weiterfahren kann. `defer` wird von go zum Beispiel bei out of bounds bei einem Array gebraucht. Auch wir können `panic` brauchen, um defensiv zu programmieren und unerwartetes bemerkbar zu machen.

Beide folgenden Funktionen werden eine `panic` werfen:

```go
func outOfBounds() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[5])
}

func myPanic() {
	panic("something bad happened")
}
```

### Wie kann man eine `panic` behandeln oder auffangen?

Um eine panic aufzufangen braucht man `defer` und `recover`.

`defer` wird bei einer Funktion am ende aufgerufen und ist besonders praktisch um verbindungen zu schliessen, zum beispiel beim schreiben von einer Datei der Datenbank. `defer` wird immer aufgerufen, auch wenn die funktion crashed. Das bedeutet, selbst wenn man ein Array out of bounds error hat, wird `defer` noch aufgeufen

Um die `panic` dann abzufagnen braucht man `recover`. Diese wird in der methode, welche mit `defer` aufgerufen wird eingebaut. Mit `recover` kann man dann die nachricht ausgeben.

```go

func safeRun() {
    defer handlePanic()
    outOfBounds()
}
func handlePanic() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from panic:", r)
    }
}
```

```go
package main

import "fmt"

func main() {
safeRun()
fmt.Println("Program continues after panic is recovered.")
}

func safeRun() {
defer handlePanic()
outOfBounds()
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
```