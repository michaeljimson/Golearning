package main

import "fmt"

func f1() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
	// for + TAB自动生成
}

func f2() {
	i := 1
	for ; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func f3() {
	m := 1
	for m <= 10 {
		fmt.Printf("m: %v\n", m)
		m++
	}
}

func f4() {
	for {
		fmt.Println("我一直在执行啊啊啊啊啊~~~~~~~~")
	}
}
func main() {
	//f1()
	//f2()
	//f3()
	//f4()
}
