/* package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func f1(a int) { // copy
	a = 100
}

func f2(s []int) {
	s[0] = 1000
}

func f3(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func f4(name string, ok bool, args ...int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("ok: %v\n", ok)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}
func main() {
	/* r := sum(1, 2)
	fmt.Printf("r: %v\n", r) */
	/* 	x := 200
	   	f1(x)
	   	fmt.Printf("x: %v\n", x) */
	/* s := []int{1, 2, 3}
	f2(s)
	fmt.Printf("s: %v\n", s) */
	/* f3(1, 2, 3, 4)
	f3(3, 4, 5, 6, 7) */
	f4("michale", true, 1, 2, 3, 4)
}
 */