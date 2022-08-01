package main

import "fmt"

/* func f2() {
	var num int
	fmt.Println("请输入一个数字：")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
} */

/* func f3() {
	if age := 20; age > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
} */

func f3() {
	var c string
	fmt.Println("请输入一个字符：")
	fmt.Scan(&c)

	if c == "M" {
		fmt.Println("Monday")
	} else if c == "T" {
		fmt.Println("请输入第二个字符：")
		fmt.Scan(&c)
		if c == "u" {
			fmt.Println("Tuesday")
		} else {
			fmt.Println("Thursday")
		}
	} else if c == "W" {
		fmt.Println("Wednesday")
	} else if c == "F" {
		fmt.Println("Friday")
	} else if c == "S" {
		fmt.Println("请输入第二个字符：")
		fmt.Scan(&c)

		if c == "a" {
			fmt.Println("Saturday")
		} else {
			fmt.Println("Sunday")
		}
	}

}
func main() {
	/* a := 1
	b := 2
	// gofmt
	if a > b {
		fmt.Println("a")
	} else {
		fmt.Println("b")
	} */

	/* flag1 := true
	if flag1 {
		fmt.Printf("flag1: %v\n", flag1)
	} */

	/* if age := 20; age > 18 {
		fmt.Println("成年")
	} else {
		fmt.Printf("未成年")
	} */

	/* a := 100
	if a { // 不能使用0或非0表示真假
		fmt.Println("真")
	} */
	/* a := 2
	b := 3
	if a < b // 一条语句也必须加大括号{}
		fmt.Println("a")
	else
		fmt.Println("b") */
	/* var name string
	var age int
	var email string
	fmt.Println("请输入name,age,email,用空格分隔")
	fmt.Scan(&name, &age, &email)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("email: %v\n", email) */
	/* f2() */
	f3()
}
