/* package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("Hello,%s", name)
}
func test(name string, f func(string)) {
	f(name)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func cal(operator string) func(int, int) int {
	switch operator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}
func main() {
	//test("jimson", sayHello)
	ff := cal("+")
	r := ff(1, 2)
	fmt.Printf("r: %v\n", r)
	ff = cal("-")
	r = ff(1, 2)
	fmt.Printf("r: %v\n", r)
}
*/