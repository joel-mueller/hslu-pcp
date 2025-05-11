package DELETEME

import (
	"fmt"
)

func Apply(f func(int) int, val int) int {
	return f(val)
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}

func square(val int) int {
	return val * val
}

func Demoo() {
	fmt.Println(Apply(double, 4))
	fmt.Println(Apply(triple, 5))
	fmt.Print(Apply(square, 4))
}
