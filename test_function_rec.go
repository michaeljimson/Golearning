/* package main

import "fmt"

func f1() {
	//fmt.Println("....")
	//f1() //死递归
	var s int = 1
	for i := 1; i <= 5; i++ {
		s *= i
	}
	fmt.Printf("s: %v\n", s)
}

func rec(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * rec(n-1)
	}
}

func f3(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return f3(n-1) + f3(n-2)
	}
}

func main() {
	// f1()
	/* r := rec(6)
	fmt.Printf("r: %v\n", r) */
	r := f3(7)
	fmt.Printf("r: %v\n", r)
}
 */