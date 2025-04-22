package main
import ("fmt")

func main() {
var arr3 = [...]string{"Volvo", "Ford", "Tesla"} // here length is inferred
fmt.Printf(arr3[2])
}
