package main

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
	/* name := "tom"
	age := "20"
	s := strings.Join([]string{name, age}, ",")
	fmt.Printf("s: %v\n", s) */
	/* 直写缓存区，效率较高
	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer.String(): %v\n", buffer.String()) */
	// 转义字符

	/* // s := "hello\tworld"
	// s := "I'm tom"
	s := "C:\\program files\\a"
	fmt.Printf("s: %v\n", s)
	*/
	/*
		// 字符串的切片操作
		str := "hello world"
		n := 2
		m := 5
		fmt.Printf("str[n]: %c\n", str[n])
		fmt.Printf("str[n:m]: %v\n", str[n:m])
		fmt.Printf("str[n:]: %v\n", str[n:])
		fmt.Printf("str[:m]: %v\n", str[:m]) */

	/* 	//字符串函数

	   	s := "Hello World"
	   	fmt.Printf("len(s): %v\n", len(s))
	   	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, " "))
	   	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))
	   	fmt.Printf("strings.ToLower(s): %v\n", strings.ToLower(s))
	   	fmt.Printf("strings.HasPrefix(\"Hello\"): %v\n", strings.HasPrefix(s, "Hello"))
	   	fmt.Printf("strings.HasSuffix(\"World\"): %v\n", strings.HasSuffix(s, "World"))

	   	fmt.Printf("strings.Index(s, \"ll\"): %v\n", strings.Index(s, "ll"))
	   	fmt.Printf("strings.LastIndex(s, \"ld\"): %v\n", strings.LastIndex(s, "ld"))
	*/
}
