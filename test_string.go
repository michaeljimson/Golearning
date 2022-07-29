package main

import (
	"fmt"
	"strings"
)

func main() {
	/* var s string = "Hello world"
	var s1 = "Hello world"
	s3 := "hello world"

	fmt.Printf("s: %v\n", s)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s3: %v\n", s3)

	s4 := `
	line1
	line2
	line3
	`
	fmt.Printf("s4: %v\n", s4) */

	// 字符串连接
	/* s1 := "tom"
	s2 := "20"
	msg := s1 + s2
	fmt.Printf("msg: %v\n", msg) */
	/* s1 := "tom"
	s2 := "20"
	msg := fmt.Sprintf("name=%s, age=%s", s1, s2)
	fmt.Printf("msg: %v\n", msg) */
	name := "tom"
	age := "20"
	s := strings.Join([]string{name, age}, ",")
	fmt.Printf("s: %v\n", s)
}
