/* package main

import "fmt"

//执行过程 初始化变量->init()->main()
var i int = initVar()

/* func init() {
	fmt.Println("init.....")
}

func init() {
	fmt.Println("init2...")
} */
func init() {
	fmt.Println("init3...")
}
func initVar() int {
	fmt.Println("initVar...")
	return 100
}
func main() {

	fmt.Println("main.....")
}
 */