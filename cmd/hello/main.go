package main

import (
	"fmt"
	"golearn/internal/hello"
	"math/cmplx"
)

// variables defined at package level (type at end)
var c, python, java bool

// variables defined with initialisers
var a, b int = 1, 2

// basic types (+ a block of variable declarations)
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 21i)
)

// numeric constants
const (
	Big = 1 << 100
	Small = Big >> 99
)
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(hello.SayHello("world"))

	fmt.Println("\nVariables:")
	// variables defined at function level
	var i int

	// variables defined without type, but inferred from init
	var javascript, golang = true, "no!"

	// short variable declations (can only be used inside functions)
	c := 3		// same as var c = 3
	fmt.Println(i, c, python, java, a, b, javascript, golang, c)

	fmt.Println("\nSome basic types:")
	fmt.Printf("ToBe Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("MaxInt Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("z Type: %T Value: %v\n", z, z)

	fmt.Println("Zero values:", hello.ZeroValues())
	fmt.Println("Type conversion:", hello.TypeConversion())
	fmt.Println("Type inference:", hello.TypeInference())

	fmt.Println("Constants:", hello.Constants())

	fmt.Println("\nNumeric constants:")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big))  // throws error as the value of 'Big' is too large

}
