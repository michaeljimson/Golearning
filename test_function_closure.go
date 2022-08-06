/* package main

import (
	"fmt"
	"strings"
)

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	/* f := add()
	r := f(10)
	fmt.Printf("r: %v\n", r)
	r = f(20)
	fmt.Printf("r: %v\n", r)
	r = f(30)
	fmt.Printf("r: %v\n", r)

	fmt.Println("--------------------")

	f1 := add()
	r = f1(100)
	fmt.Printf("r: %v\n", r)
	r = f1(100)
	fmt.Printf("r: %v\n", r) */
	/* jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(txtFunc("test")) */
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))

	f1, f2 = calc(100)
	fmt.Println(f1(10), f2(20))
	fmt.Println(f1(30), f2(40))
	fmt.Println(f1(50), f2(60))
}
 */