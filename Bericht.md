## Wieso wurde Go Entwickelt?

Go wurde ab 2007 bei Google von Robert Griesemer, Rob Pike und Ken Thompson entwickelt und 2012 veröffentlicht. Die Motivation war, eine Sprache zu entwickeln, die so schnell wie C, so lesbar wie Python und gut für Nebenläufigkeit ist. Viele Go-Entwickler waren mit C++ unzufrieden. Weitere Informationen dazu gibt es auf [Wikipedia Go](https://en.wikipedia.org/wiki/Go_(programming_language)).

## Structural & Nominal Typing

Go verwendet sowohl strukturelle als auch nominale Typisierung. Die Code Beispiele aus diesem Kapitel stammen aus **Datastructures**.

### Nominal Typing

Beim Nominal Typing werden zwei Strukturen oder Klassen nicht als gleich oder kompatibel angesehen, auch wenn sie die gleichen Datentypen und Methoden haben.
In Go gilt dies für alle primitiven Datentypen und Strukturen. Einer Funktion, die eine Struktur oder einen Zeiger auf eine Struktur als Parameter hat, muss genau eine solche Struktur übergeben werden. Eine andere Struktur, die aber den gleichen Methodennamen und die gleichen Attribute hat, kann dieser Funktion nicht übergeben werden.

Wir haben hier zwei Stacks. Einer hat intern ein Slice und der andere eine Linked List.

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

Dies ist eine Funktion, die Statistiken über den Stack ausgibt. Die Methode soll für den Stack funktionieren.

```go
func GetStatsStack[T any](stack *Stack[T]) string {
    if stack.Empty() {
        return "The stack is empty"
    }
    return fmt.Sprintf("The size is %d", stack.Size())
}
```

Diese Funktion kann Statistiken aus dem Stack ausgeben, aber nicht aus der StackList, da es sich um einen anderen Typ handelt.

```go
stack := Stack[int]{}
stackList := StackList[int]{}
fmt.Println(GetStatsStack(&stack))
fmt.Println(GetStatsStack(&stackList)) // error, StackList ist nicht vom Typ Stack
```

### Structural Typing

Neben Nominal Typing verwendet Go auch Structural Typing. Structural Typing wird in Go ausschliesslich für Interfaces verwendet. Wenn eine Struktur in Go die gleichen Methoden wie ein Interface hat, hat die Struktur auch automatisch das Interface implementiert.
Das folgende Beispiel ist ein Interface für Datenstrukturen.

```go
type Datastructures interface {
    Empty() bool
    Size() int
}
```

Da sowohl `Stack` als auch `StackList` diese beiden Funktionen besitzen, implementieren sie automatisch das Interface `Datastructure`.
Das folgende Beispiel ist eine Funktion, die keines der beiden Stacks als Parameter annimmt, sondern das `Datastructure` Interface. Die Methoden des `Datastructure` Interface wurden sowohl von `Stack` als auch von `StackList` implementiert.

```go
func GetStats(datastructures Datastructures) string {
    if datastructures.Empty() {
        return "The datastructures is empty"
    }
    return fmt.Sprintf("The size is %d", datastructures.Size())
}
```

Der Funktion kann sowohl ein `Stack` als auch eine `StackList` übergeben werden.

```go
fmt.Println(GetStats(&stackList))
fmt.Println(GetStats(&stack))
```

Structural Typing hat den Vorteil, dass man die `structs` externer Packages einfach mit `interfaces` mocken und auch gut testen kann. Auch Entwurfsmuster wie [Adapter](https://refactoring.guru/design-patterns/adapter) können vereinfacht oder sogar ganz weggelassen werden. Die Gefahr kann aber auch sein, dass wenn man ein Interface oder eine Implementation ändert, zwei Dinge nicht mehr kompatibel sind und ein Fehler auftritt.

## `defer`, `panic` und `recover`

### Was ist `defer`?

In Go wird `defer` immer am Ende einer Funktion aufgerufen und ist besonders nützlich, um Verbindungen zu schliessen, z.B. beim Schreiben in eine Datei oder in eine Datenbank. `defer` wird immer aufgerufen, auch wenn die Funktion eine Panic oder einen Error hat. D.h. selbst wenn man einen Array out of bounds error hat, wird `defer` immer noch aufgerufen.

Hier ein Beispiel, wie `defer` in **LanguageDetection** verwendet wurde.

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

### Wann braucht man `panic` ?

In Go wird `panic` benutzt, wenn etwas grosses Unerwartetes passiert und das Programm nicht weiterlaufen kann. `panic` wird von Go z.B. bei out of bounds in einem Array gebraucht. Wir können `panic` auch benutzen, um defensiv zu programmieren und unerwartete Dinge anzuzeigen.
Die folgende Methode liest den obersten Wert aus dem Stack und löst eine Panic aus, wenn der Stack leer ist. Die stammt aus **Datastructures**.

```go
func (stack *Stack[T]) Peek() T {
    if len(stack.stack) == 0 {
        panic("stack is empty")
    }
    return stack.stack[len(stack.stack)-1]
}
```

### Wie kann man eine `panic` behandeln oder auffangen?

Um die `panic` wieder aufzufangen, benötigt man `recover` und `defer`. Wenn man eine `panic` erwartet kann eine funktion mit `defer` danach aufgerufen werden. Diese kann dann mit `recover` den error auffangen und ausgeben. Folgendes Beispiel kommt aus **Panic**.

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

Wenn eine Goroutine einen Wert in eine Variable schreibt und eine andere Goroutine diese Variable später lesen möchte, legt das Go Memory Modell fest, wann und ob der Leser garantiert den neuen oder noch den alten Wert liest. Um diese Garantie zu gewährleisten, gibt es folgende Möglichkeiten:
- Channels
- Ein Mutex (`sync.Mutex`)
- Atomic Variablen (`sync.Atomic`)

Waiting Groups (`sync.WaitGroup`) sind sehr praktisch, um auf mehrere offene Goroutinen zu warten, und daher eine Erwähnung wert :)

Die Beispiele in diesem Kapitel kommen aus **Bank**.

### Synchronisation mit Mutex

Im Beispiel Bank wurde in der Klasse ein Mutex definiert. Dieser wird für den Zugriff auf den Saldo des Bankkontos verwendet.

```go
type Account struct {
    balance          int
    mu               sync.Mutex
    transactionCount uint64
}
```

Um eine Transaktion durchzuführen, muss das Lock geöffnet und wieder geschlossen werden. Dies geschieht am besten mit `defer`.

```go
func (acc *Account) Deposit(amount int) {
    acc.mu.Lock()
    defer acc.mu.Unlock()
    acc.balance += amount
    atomic.AddUint64(&acc.transactionCount, 1)
}
```

### Synchronisatzion mit Atomic Variablen

In Go ist es möglich, eine Variable in eine atomare Variable umzuwandeln. Eine Atomic Variable wird wie eine normale Variable deklariert. In der Struktur Account wird sie als `transactionCount uint64` gespeichert. Diese Variable kann dann mit `atomic.AddUint64(&acc.transactionCount, 1)` sicher um eins erhöht werden.

### Waiting Groups

Die Waiting Group ist sehr nützlich, um auf Routinen zu warten. In der `Demo()` Funktion der Bank wurde eine Waiting Group mit `var wg sync.WaitGroup` erstellt. Danach wurde mit `wg.Add(2)` auf zwei Funktionen gewartet. Wenn eine Funktion fertig ist, kann sie mit `defer wg.Done()` als fertig markiert werden. Am Ende der Funktionen kann mit `wg.Wait()` gewartet werden.

Weitere Informationen zum Go Memory Modell gibts auf [The Go Memory Model] (https://go.dev/ref/mem).

## `goroutines`, `channels` & `select`

Die Code Beispiele in diesem Kapitel kommen aus **Routines**.

### `goroutine`

Eine Go-Routine beschreibt einen leichtgewichtigen Thread, der von der Go-Runtime verwaltet wird. Jede Funktion kann einfach mit `go functionname()` als Go-Routine aufgerufen werden. Hier ist ein Beispiel von [Goroutine Tour of Go](https://go.dev/tour/concurrency/1).

### `channels`

Channels sind typisierte Leitungen, über die Daten gesendet und empfangen werden können. Ein Channel kann einfach erstellt werden mit `c := make(chan int)`. Danach kann ein Channel als Parameter an eine Routine übergeben werden: `go LongLastingTask(c)`. Werte werden mit `c <- value` in Channels geschrieben und mit `value <- c` geholt. Hier ist ein Beispiel aus den Übungen.

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

Hier ein Beispiel von [Channels Tour of Go](https://go.dev/tour/concurrency/2).

### `select`

Die Select-Anweisung lässt eine Routine auf mehrere Kommunikationsoperationen warten. `select` blockiert, bis einer seiner `case`'s ausführbar ist, und führt ihn dann aus. Sind mehrere gleichzeitig ausführbar, wird zufällig ausgewählt. Hier ein Beispiel von [Select Tour of go] (https://go.dev/tour/concurrency/6) und hier ein Beispiel, wie `select` in der Übung verwendet wurde.

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

- Joel: Ich finde, dass Go eine tolle Sprache ist. Während meines Studiums habe ich viel Code in Java und Python geschrieben. Go hat das Beste aus beiden Welten, was mich sehr motiviert hat, die Sprache zu lernen. Das Einzige, was ich nicht so gut finde, ist das Structural Typing von Interfaces. Es ist mir ein bisschen zu offen und ich denke, es kann auch zu grösseren Fehlern führen, wenn man nicht aufpasst. Ich frage mich auch, wie es ist, ein grosses Projekt mit Go zu machen, aber das werde ich sicher in der Zukunft machen.
- Leo: TODO (je min. 1 Abschnitt pro Team-Mitglied)

## Gebrauchte Ressourcen

- [W3 Schools Go](https://www.w3schools.com/go/go_getting_started.php)
- [Go by example](https://gobyexample.com/)
- [Tour of go](https://go.dev/tour/welcome/1)
- [Go Dokumentation](https://go.dev/doc/)

## Übersicht Codebeispiele

- **AdventOfCode**: Programmierübung zu Clojure Woche 2, 5. Aufgabe
- **Bank**: Beispiel Transaktionen von einem Bankkonto für das Go Memeory Modell
- **Datastructures**: PCP-Übung Woche 1: Einstieg - C & Java revisited, 2. Aufgabe: ADT Stack in C (Array-Implementierung)
- **LanguageDetection**: Programmierübung zu Clojure Woche 2, 7. Aufgabe
- **Panic**: Beispiel auffangen von `panic`
- **Routines**: PCP-Übung zu Java 8, Teil 3 + 4 (Woche 9), 3. CompletableFuture mit zwei nebenläufigen Tasks