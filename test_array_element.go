package main

import "fmt"

func test1() {
	var a1 = [2]int{100, 200}
	fmt.Printf("a1: %v\n", a1)
	fmt.Println("------------")
	a1[0] = 200
	a1[1] = 100
	fmt.Printf("a1: %v\n", a1)

	/* // 数组长度越界
	a1[3] = 1000 */
}

func test2() {
	// 数组的长度
	var a1 = [3]int{1, 2, 3}
	fmt.Printf("len(a1): %v\n", len(a1))

	var a2 = [...]int{1, 2, 3, 4}
	fmt.Printf("len(a2): %v\n", len(a2))
}

func test3() {
	// 数组的遍历 1. 根据长度和下标

	var a1 = [3]int{1, 2, 3}
	for i := 0; i < len(a1); i++ {
		fmt.Printf("a1[%v]: %v\n", i, a1[i])
	}
}

func test4() {
	var a1 = [3]int{1, 2, 3}
	/* for i, v := range a1 {
		fmt.Printf("a1[%v]: %v\n", i, v)
	} */
	for _, v := range a1 {
		fmt.Printf("v: %v\n", v)
	}
}
func main() {
	// test1()
	// test2()
	// test3()
	test4()
}
