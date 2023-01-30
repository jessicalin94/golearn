package hello

import (
	"fmt"
	"math"
)

// constants
const Pi = 3.14

func SayHello(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
func ZeroValues() string {
	// Zero values when no initial value given on declaration
	var i int		// numeric = 0
	var f float64	// 0
	var b bool		// false
	var s string	// ""
	return fmt.Sprintf("%v %v %v %q\n", i, f, b, s)
}

func TypeConversion() string {
	// Type conversion
	// T(v) converts the value to the type T 
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	return fmt.Sprintln(x, y, z)
}

func TypeInference() string {
	// When declaring an untyped numeric constant, it'll be inferred by the precision of the value declared
	v := 42.1i
	return fmt.Sprintf("v is of the type %T\n", v)
}

func Constants() string {
	// declared using the 'const' keyword
	// can be character, string, boolean, numeric
	// cannot use shorthand to declare constants ( := )
	const World = "世界"
	return fmt.Sprintln("Hello", World, "Happy", Pi, "Day")
	return fmt.Sprintln("Happy", Pi, "Day")
}