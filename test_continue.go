package main

import "fmt"

func test1() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("i: %v\n", i)
		} else {
			continue
		}
	}
}

func test2() {
MYLABEL:

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 2 && j == 2 {
				continue MYLABEL
			}
			fmt.Printf("%v:%v\n", i, j)
		}
	}
}

func main() {
	//test1()
	//test2()
}
