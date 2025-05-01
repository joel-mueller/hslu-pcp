# Fokuspunkte

- Defer, panic und recover
- Goroutines, Channels & Select
- Structural & Nominal Typing
- The Go Memory Model

Abgabe: Ihre gewählten 3 bis 7 Fokuspunkte, jeweils
inkl. kurze Erklärung (stichwortartig, ggf. inkl.kurze
Pseudo-Code-Sequenzen)

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