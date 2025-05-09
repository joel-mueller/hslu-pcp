## Wieso wurde go Entwickelt?

Go wurde ab 2007 bei Google von Robert Griesemer, Rob Pike, und Ken Thompson entwickelt und 2012 veröffentlicht. Die Motivation war es, eine sprache zu entwickeln, welche so schnell ist wie C, so lesbar wie Python und gut ist für nebenläufigkeit. Viele Entwickler von Go waren unglücklich von C++. Mehr informationen gibts auf [Wikipedia Go](https://en.wikipedia.org/wiki/Go_(programming_language)).

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

Dies hat der Vorteil, das man `structs` aus externen Packages einfach mit `interfaces` Mocken und auch gut Testen kann. Auch Entwurfsmuster wie [Adapter](https://refactoring.guru/design-patterns/adapter) können vereinfacht oder sogar komplett weggelassen werden. Die gefahr kann aber auch sein, das wenn man ein Interface oder eine Implementation ändert, zwei dinge nicht mehr Kompatibel sind und es zu einem Fehler kommt.

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

## The Go Memory Model

https://go.dev/ref/mem

Wenn eine Goroutine einen Wert in eine Variable schreibt und eine andere Goroutine später diese Variable lesen will, legt das Go Memory modell fest, wann und ob der leser garantiert die Neue liest oder noch die alte. Um diese Garantie zu gewährleisten, gibt es folgende möglichkeiten:

- Channels
- Ein Mutex (`sync.Mutex`)
- Atomic Variablen (`sync.Atomic`)

Waiting groups (`sync.WaitGroup`) sind sehr praktisch, um auf mehrere offene Goroutinen warten, und demnach auf eine Erwähnung wert :)

### Synchronisation mit Mutex

In dem Bank beispiel wurde ein mutex in der Klasse definiert. Dieser soll für den Zugriff auf die Balance vom Bank Account dienen.

```go
type Account struct {
    balance          int
    mu               sync.Mutex
    transactionCount uint64
}
```

Um eine Transaktion muss der Lock geholt werden und danach auch wieder geschlossen werden. Das macht man am besten mit `defer`.

```go
func (acc *Account) Deposit(amount int) {
    acc.mu.Lock()
    defer acc.mu.Unlock()
    acc.balance += amount
    atomic.AddUint64(&acc.transactionCount, 1)
}
```

### Synchronisatzion mit Atomic Variablen

Auch kann man aus einer Variable eine Atomic variable machen. Eine Atomic variable wird als normale variable deklariert. In der Account struct wurde sie als `transactionCount uint64` gespeichert. Diese Variable kann dann mit `atomic.AddUint64(&acc.transactionCount, 1)` um ein erhöht werden.


### Waiting Groups

Die Waiting Group ist sehr praktisch, um auf Goroutines zu warten. In der `Demo()` funktion von der bank wurde mit `var wg sync.WaitGroup` eine Wiaint group erstellt. Danach wurde mit
`wg.Add(2)` der Auftrag gegeben auf Zwei funktionen zu warten. Wenn eine funktion fertig ist, kann sie mit `defer wg.Done()` als fertig gekennzeichnet werden. Am ende der Funktionen kann dann mit `wg.Wait()` gewartet werden. Eine Waiting Group ist mehr gedacht, um auf goroutines zu warten, kann aber auch gebraucht werden um variablen zu synchronisieren.

## `goroutines`, `channels` & `select`

### `goroutine`

Eine goroutine beschreibt einen leichtgewichtigen thread welcher von der Go Runtime gemanaged wird. Jede funktion kann ganz einfach mit `go functionname()` als go routine aufgerufen werden. Hier ein Beispiel von [Goroutine Tour of Go](https://go.dev/tour/concurrency/1).

### `channels`

Channels sind eine typisierte Leitung durch welche man Daten senden und bekommen kann. Ein channel kann ganz einfach mit `c := make(chan int)` erstellt werden. Danach kann ein channel einer Goroutine als parameter mitgegeben werden: `go LongLastingTask(c)`. Channels werden mit `c <- value` geschrieben und mit `value <- c` geholt. Hier ein beispiel aus den übungen. Auch gibt es ein Beispiel von [Channels Tour of Go](https://go.dev/tour/concurrency/2)

```go
func LongLastingTask(c chan int) {
    DoBlockingWait(3000)
    fmt.Print("3000")
    c <- 3000
}
```

```go
c := make(chan int)
go LongLastingTask(c)
go EvenLongerLastingTask(c)
s, t := <-c, <-c
```


### `select`

Das select statement lässt eine goroutine auf mehrere kommunikations operationen warten. Ein select blockiert solange bis einer seiner cases ausführbar ist, dann führt es diesen aus. Falls mehrere gleichzeitig ausführbar sind wird zufällig ausgewählt. Hier ein beispiel von [Select Tour of go](https://go.dev/tour/concurrency/6) und hier ein Beispiel wie wir es gebraucht haben in Routines Channels

```go
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
```

## Technisches Team-Fazit

TODO
- Speicher Grosse von runntime file, wie ist das?
- Grosse standard libary

## Persönliches Fazit

- Joel: Ich finde Go eine sehr tolle sprache. Ich habe im Studium viel code in Java und Python geschrieben. Go hat von beiden welten das beste drin, was mich sehr motivierte die sprache zu lernen. Das einzige was ich nicht so gut finde ist das Structural Typing von Interfaces. Dies ist mir etwas zu offen und ich denke es kann auch zu grösseren fehler füren wenn man nicht gut aufpasst. Ich frage mich auch wie es ist ein Grosses Projekt mit Go umzusetzen, was ich aber in Zukunft sicher machen will.
- Leo: TODO (je min. 1 Abschnitt pro Team-Mitglied)

## Gebrauchte Ressourcen

- [W3 Schools Go](https://www.w3schools.com/go/go_getting_started.php)
- [Go by example](https://gobyexample.com/)
- [Tour of go](https://go.dev/tour/welcome/1)
- [Go Dokumentation](https://go.dev/doc/)