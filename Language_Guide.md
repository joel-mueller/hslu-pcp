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

## Arrays

### Creating and array

```go
arr1 := [6]int{10, 11, 12, 13, 14,15}
arr2 := [5]int{4,5,6,7,8} // if not every value is initialized, the default value is set
var arr3 = [...]string{"Volvo", "Ford", "Tesla"} // here length is inferred, but when array is defined, the size cant be changed
arr4 := [5]int{1:10,2:40} // initalize the second and third value, the rest to default value
```

### Aarray operations

```go
fmt.Printf(arr2)
fmt.Printf(arr3[2])
arr3[2] = "Volkswagen" // change element
fmt.Printf(len(arr4))
```

### Characteristics

- Copy is by value
- Fixed size

## Slices

### Creating a slice

```go
myslice1 := []int{}
myslice2 := []int{1,2,3}
myslice := arr1[2:4] // creating a slice from an array
myslice1 := make([]int, 5, 10) // make([]type, length, capacity), if capazity is not defined, its the same as lenght
```

### Slice operations

- `len()` gives the number of elements in the slice back
- `cap()` gives the capacitry of the slice back (the number it can grow or shrink to)

```go
myslice1 = append(myslice1, 20, 21) // append(slice_name, element1, element2, ...)
myslice3 = append(myslice1, myslice2...) // append(slice1, slice2...)
```

### Copy function usage

```go
neededNumbers := numbers[:len(numbers)-10]
numbersCopy := make([]int, len(neededNumbers))
copy(numbersCopy, neededNumbers)
```

### Characteristics

- Copy is by refference
- Size is flexible and can grow and shrink

## Operators

- [Arithmetic Operators](https://www.w3schools.com/go/go_arithmetic_operators.php)
- [Assignment Operators](https://www.w3schools.com/go/go_assignment_operators.php)
- [Comparisson Operators](https://www.w3schools.com/go/go_comparison_operators.php)
- [Logical Operators](https://www.w3schools.com/go/go_logical_operators.php)
- [Bitwise operators](https://www.w3schools.com/go/go_bitwise_operators.php)

## Condition

```go
time := 22
if time < 10 {
  fmt.Println("Good morning.")
} else if time < 20 {
  fmt.Println("Good day.")
} else {
  fmt.Println("Good evening.")
}
```

## Switch Statement

```go
switch day {
case 1,3,5:
  fmt.Println("Odd weekday")
case 2,4:
  fmt.Println("Even weekday")
case 6,7:
  fmt.Println("Weekend")
default:
  fmt.Println("Invalid day of day number")
}
```

## For Loop

```go
for i:=0; i <= 100; i+=10 {
  fmt.Println(i)
}
```

- The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
- The break statement is used to break/terminate the loop execution.

### Range

```go
for idx, val := range fruits { // index, value := range array|slice|map
   fmt.Printf("%v\t%v\n", idx, val)
}
```

## Functions

```go
func myFunction(x int, y int) int {
  return x + y
}
```

### Named return values

```go
func myFunction(x int, y int) (result int) {
  result = x + y
  return
}
```

### Multiple return values

```go
func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  a, b := myFunction(5, "Hello")
  fmt.Println(a, b)
}
```

## Go Structs

```shell
package main
import ("fmt")

type Person struct {
  name string
  age int
  job string
  salary int
}

func main() {
  var pers1 Person
  var pers2 Person

  // Pers1 specification
  pers1.name = "Hege"
  pers1.age = 45
  pers1.job = "Teacher"
  pers1.salary = 6000

  // Pers2 specification
  pers2.name = "Cecilie"
  pers2.age = 24
  pers2.job = "Marketing"
  pers2.salary = 4500

  // Print Pers1 info by calling a function
  printPerson(pers1)

  // Print Pers2 info by calling a function
  printPerson(pers2)
}

func printPerson(pers Person) {
  fmt.Println("Name: ", pers.name)
  fmt.Println("Age: ", pers.age)
  fmt.Println("Job: ", pers.job)
  fmt.Println("Salary: ", pers.salary)
}
```

## Maps

Make maps with `var` and `:=`

```shell
var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
fmt.Printf("a\t%v\n", a)
fmt.Printf("b\t%v\n", b)
```

Using the `make()` function

```shell
var a = make(map[string]string) // The map is empty now
a["brand"] = "Ford"
a["model"] = "Mustang"
a["year"] = "1964"
```

> Note: Always use `make()` to create an empty map. Doing it in a different way will end up in a runtime panic

### Allowed Types in Map

As key is everything allowed except: (this is because the equality oparator for comparing with `==` is there not defined)
- Slices
- Maps
- Functions

As value is everything allowed

### Operations in Map

```shell
fmt.Printf(a["brand"]) // get an entry
a["year"] = "1970" // update an entry
delete(a,"year") // deleting an entry
val3, ok3 := a["day"] // Checking for existing key and its value
_, ok4 := a["model"] // Only checking for existing key and not its value

// Iterate over maps
for k, v := range a {
    fmt.Printf("%v : %v, ", k, v)
}
```
