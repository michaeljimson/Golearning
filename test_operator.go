package main

import "fmt"

func getA() int {
	return 100
}
func main() {
	/* a := 100
	b := 20
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	x := a + b
	fmt.Printf("x: %v\n", x)

	x = a - b
	fmt.Printf("x: %v\n", x)

	x = a * b
	fmt.Printf("x: %v\n", x)

	x = a / b
	fmt.Printf("x: %v\n", x)

	x = a % b
	fmt.Printf("x: %v\n", x)
	*/

	/* a := 100
	a++
	fmt.Printf("a: %v\n", a)
	a--
	fmt.Printf("a: %v\n", a) */

	/* a := 10
	b := 5
	r := a == b
	fmt.Printf("r: %v\n", r)
	r = a > b
	fmt.Printf("r: %v\n", r)
	r = a < b
	fmt.Printf("r: %v\n", r)
	r = a >= b
	fmt.Printf("r: %v\n", r)
	r = a <= b
	fmt.Printf("r: %v\n", r)
	r = a != b
	fmt.Printf("r: %v\n", r) */

	/* a := true
	b := false

	r := a && b
	fmt.Printf("r: %v\n", r)

	r = a || b
	fmt.Printf("r: %v\n", r)

	fmt.Printf("a: %v\n", !a) */

	/* a := 4
	fmt.Printf("a: %b\n", a) //0100
	b := 8
	fmt.Printf("b: %b\n", b) //1000

	r := a & b
	fmt.Printf("r: %b\n", r)

	r = a | b
	fmt.Printf("r: %v\n", r)

	r = a ^ b
	fmt.Printf("r: %v\n", r)

	a = a << 2
	fmt.Printf("a: %v\n", a)

	b = b >> 2
	fmt.Printf("b: %v\n", b) */

	a := 100
	a = 200
	fmt.Printf("a: %v\n", a)
	a = getA()
	fmt.Printf("a: %v\n", a)
	a += 100
	fmt.Printf("a: %v\n", a)
}
