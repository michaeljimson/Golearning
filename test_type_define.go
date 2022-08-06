/* package main

import "fmt"

func main() { 
	/*
		//类型定义
		type MyInt int
		var i MyInt
		i = 100
		fmt.Printf("i: %v\n", i)
		fmt.Printf("i: %T\n", i)
	*/
	// 类型别名
	type MyInt = int
	var i MyInt
	i = 100
	fmt.Printf("%v: %T\n", i, i)
}
*/