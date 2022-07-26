/* package main

import "fmt"

func test1() {
	// 类型的声明
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m1: %T\n", m1)
}
func test2() {
	var m1 = map[string]string{"name": "jimson", "age": "26", "email": "897638991@qq.com"}
	fmt.Printf("m1: %v\n", m1)

	m2 := make(map[string]string)
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["email"] = "897638991@qq.com"
	fmt.Printf("m2: %v\n", m2)
}

func test3() {
	m1 := map[string]string{"name": "jimson", "age": "26", "email": "897638991@qq.com"}
	var key = "name"

	var value = m1[key]
	fmt.Printf("value: %v\n", value)
}

func test4() {
	var m1 = map[string]string{"name": "jimson", "age": "26", "email": "897638991@qq.com"}
	var k1 = "name"
	var k2 = "age1"
	v, ok := m1[k1]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	fmt.Println("----------")
	v, ok = m1[k2]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)

}
func main() {
	// test1()
	// test2()
	// test3()
	test4()
}
*/