/* package main

import "fmt"

func test1() {
	fmt.Println("我既没有参数，也没有返回值。。。")
}

func test2(a int, b int) int {
	return a + b
}

func sum3() (name string, age int) {
	/* name = "michael"
	age = 20 */
	n := "jimson"
	a := 20
	//return name, age
	return n, a
}
func main() {
	//test1()
	name, age := sum3()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)

}
 */