# Language Guide

```go
```

## Variable Declaration

```go
var student1 string = "John" //type is string
var student2 = "Jane" //type is inferred
x := 2 //type is inferred, only in functions
var a string // initializes default value
var a, b, c, d int = 1, 3, 5, 7 // multiple assignement
var (
     a int
     b int = 1
     c string = "hello"
   ) // declaration in block, works also with constants
const CONSTNAME type = value // no default values
```

## Print

```go
fmt.Println(HELLO, WORLD, "\n") // multiple prints
var i string = "Hello"
fmt.Printf("i has value: %v and type: %T\n", i, i) // %T and %v for value and type
```
> Note: There are a lot more, see also [Formatting Verbs Go](https://www.w3schools.com/go/go_formatting_verbs.php)

## Go Datatypes

- bool
- int, int8, int16, int32, int64
- float, float32, float64
