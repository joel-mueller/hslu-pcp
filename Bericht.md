# Bericht Go

Grober Themenfokus pro Sprache…
Sie bearbeiten 3 bis 7 interessante Sprachkonstrukte
oder -konzepte
– Auf welche Konstrukte und Konzepte "Ihrer" Sprache
gehen Sie im Detail ein?
– Unser Vorschlag pro Sprache siehe Git-Repo:


Qualität vor Quantität!
– Wenige Seiten (ca. 2-4) reichen durchaus, max. 5 (bei mehr
gibt's tendenziell Abzug), Inhaltsverzeichnis nicht nötig
– Fokus auf wichtige/ interessante/ spezielle SprachEigenschaften! (Was anders als bei Java?!...)
§ Erwarteter Inhalt (ergänzend zu den Folien):
– Falls interessant/relevant kurze Infos zu Vision,
Geschichte & Verbreitung
– Hauptteil: Die Sprache vorstellen (Ihre 3 bis 7
Fokuspunkte, inkl. Verweise auf Ihren Demo-Code)
– Ihr technisches Team-Fazit
– Persönliches Fazit (je min. 1 Abschnitt pro Team-Mitglied)

# Fokuspunkte

- Defer, panic und recover
- Goroutines, Channels & Select
- Structural & Nominal Typing
- The Go Memory Model

Abgabe: Ihre gewählten 3 bis 7 Fokuspunkte, jeweils
inkl. kurze Erklärung (stichwortartig, ggf. inkl.kurze
Pseudo-Code-Sequenzen)


## Structural & Nominal Typing

Go verwendet sowohl Structural, als auch Nominal Typing.

### Nominal Typing

Bei Nominal Typing werden zwei strukturen oder klassen nicht als gleich oder Kompatibel angeseehen, auch wenn sie die gleichen datentypen und Methoden haben.
In der sprache `go` ist das bei allen primitiven datentypen und structs so. Das bedeutet, wenn ich in go eine funktion habe, welche eine struct oder ein pointer zu einer struct nimmt, dann muss ich genau eine solche struct übergeben nicht eine andere struct, welche aber die gleichen namen hat und die selben attributten.

Wir haben hier zwei stacks. Eines hat intern ein slice und das andere eine Linked list. 

```go
type Stack[T any] struct {
    stack []T
}
```

```go
type StackList[T any] struct {
    head *Element[T]
}
```

Ich schreibe eine Methode, welche die statistik über das stack ausgibt. Die Methode soll für das `Stack` funktionieren. 

```go
func GetStatsStack[T any](stack *Stack[T]) string {
    if stack.Empty() {
        return "The stack is empty"
    }
    return fmt.Sprintf("The size is %d", stack.Size())
}
```

Ich kann die Statistik vom Stack ausgeben, aber nicht vom StackList, das es nicht der gleiche Typ ist.

```go
stack := Stack[int]{}
stackList := StackList[int]{}
fmt.Println(GetStatsStack(&stack))
fmt.Println(GetStatsStack(&stackList)) // error, StackList ist nicht vom Typ Stack
```

### Structural Typing

Go neben Nominal Typing, verwendet go auch structural Typing. Structural Typing wird in go aussschliesslich bei interfaces verwendet. Wenn eine struct in go die gleichen methoden, wie ein interface hat, hat die struct auch das interface.
Wir schreiben also ein Interface, welche für Datenstrukturen gedacht ist.

```go
type Datastructures interface {
    Empty() bool
    Size() int
}
```

Da sowohl `Stack` als auch `StackList` diese beiden funktionen haben, implementieren sie automaisch das Interface `Datastructure`.
Wir schreiben also folgende Funktion, welche keines der beiden Stacks nimmt, sondern ein interface, welches die methoden besitzt, welche sowhol `Stack` als auch `StackList` implementiert hat.

```go
func GetStats(datastructures Datastructures) string {
    if datastructures.Empty() {
        return "The datastructures is empty"
    }
    return fmt.Sprintf("The size is %d", datastructures.Size())
}
```

Die funktion kann jetyt für beide stacks benutzt werden. Siehe der Code auch im Ordner Datastructures.

```go
fmt.Println(GetStats(&stackList))
fmt.Println(GetStats(&stack))
```

## The Go Memory Model

### Goroutines and Concurrency

- Go verwendet *goroutines* für leichtgewichtige parallelität.
- Speicher Zugriff zwischen *goroutines* muss synchronisiert werden um race conditions zu veremeiden.

### Happens-before Relationship

- Das Memory Model definiert welche operationen garantiert sichtbar für andere *goroutinen* ist.
- Falls eine Aktion vor einer anderen passiert, dann wird garantiert das der nächste diese sieht.

### Atomic Operations
- Das ```sync/atomic``` package bietet low-level atomaren memory Zugriff mit garantierter visibility.

### Compiler and CPU Reordering
- Das Model erlaubt Compilers und CPUs Instruktionen um zu ordnen solange das happens-before Regel respektiert.

### Data Races are Bugs
- Go behandelt Data races als programmier Errors.



## `goroutines`, `channels` & `select`
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

### Was ist `defer` ?

In go wird `defer` bei einer Funktion am ende aufgerufen und ist besonders praktisch um verbindungen zu schliessen, zum beispiel beim schreiben von einer Datei der Datenbank. `defer` wird immer aufgerufen, auch wenn die funktion eine panic oder ein error hat. Das bedeutet, selbst wenn man ein Array out of bounds error hat, wird `defer` noch aufgeufen

Hier ein Beispiel wie wir defer gebraucht haben in der Language Detection:

```go
func ReadFile(path string) string {
    file, err := os.Open(path)
    if err != nil {
        log.Fatal("Error opening file:", err)
    }
    defer func(file *os.File) {
        err := file.Close()
        if err != nil {
            log.Fatal("Error Occurred while trying to clone the file", err)
        }
    }(file)
    // lesen vom file
```

Wir haben hier eine `defer` funktion definiert, welche sicher ausgeführt wird, sobald ein file da ist und die Methoden fertig ist. Dies ist besonders praktisch, da man die dinge schliessen kann, direkt dan dem man sie geöffnet hat. So hat man besonders leserlichen code und vergisst nichts.

### Wann braucht man `panic` ?

In go wird `panic` gebraucht wenn etwas grosses unerwartetes vorliegt und das Programm nicht weiterfahren kann. `panic` wird von go zum Beispiel bei out of bounds bei einem Array gebraucht. Auch wir können `panic` brauchen, um defensiv zu programmieren und unerwartetes bemerkbar zu machen.
Folgendes beispiel kommt aus dem Stack und wirft eine Panic wenn das Stack leer ist, und das oberste element gelesen werden will.

```go
func (stack *Stack[T]) Peek() T {
    if len(stack.stack) == 0 {
        panic("stack is empty")
    }
    return stack.stack[len(stack.stack)-1]
}

```

### Wie kann man eine `panic` behandeln oder auffangen?

Um die `panic` dann abzufagnen braucht man `recover` und `defer`. Diese wird in der methode, welche mit `defer` aufgerufen wird eingebaut. Mit `recover` kann man dann die nachricht dann ausgeben.

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
