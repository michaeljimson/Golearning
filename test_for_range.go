package main

import "fmt"

func f1() {
	var a = [...]int{1, 2, 3}
	/* for i, v := range a {
		fmt.Printf("%v:%v\n", i, v)
		//fmt.Printf("v: %v\n", v)
	} */
	for _, v := range a { //forr
		fmt.Printf("v: %v\n", v)
	}
}

func f2() {
	//切片
	var s = []int{1, 2, 3}
	for _, v := range s {
		fmt.Printf("v: %v\n", v)
	}
}

func f3() {
	//key : value
	m := make(map[string]string, 0)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "897638991@qq.com"

	for key, value := range m {
		fmt.Printf("%v:%v\n", key, value)
	}
}
func f4() {
	s := "hello world"
	for _, v := range s {
		fmt.Printf("v: %c\n", v)
	}
}
func main() {
	//f1()
	//f2()
	//f3()
	f4()
}
