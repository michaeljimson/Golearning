/* package main

import "fmt"

func test1() { //add
	var s1 = []int{}
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	s1 = append(s1, 300)
	fmt.Printf("s1: %v\n", s1)
}

func test2() { //del
	var s1 = []int{1, 2, 3, 4}

	s1 = append(s1[:2], s1[3:]...)
	// a = append(a[:index], a[index+1:]...)

	fmt.Printf("s1: %v\n", s1)
}

func test3() { //update
	var s1 = []int{1, 2, 3, 4}
	s1[1] = 100
	fmt.Printf("s1: %v\n", s1)
}

func test4() { //query
	var s1 = []int{1, 2, 3, 4}
	var key = 2
	for i, v := range s1 {
		if v == key {
			fmt.Printf("i: %v\n", i)
		}

	}
}

func test5() {
	var s1 = []int{1, 2, 3, 4}
	// var s2 = s1
	var s2 = make([]int, 4) //make分配内存
	copy(s2, s1)
	s2[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)

}
func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	test5()
}
*/