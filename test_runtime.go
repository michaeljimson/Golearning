/* package main

import (
	"fmt"
	"runtime"
	"time"
)

func show1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("a: %v\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}
func show2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("b: %v\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	go show1()
	go show2()

	time.Sleep(time.Second)
} */

/* package main

import (
	"fmt"
	"runtime"
	"time"
)

func show() {s
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			runtime.Goexit()
		}
	}
}
func main() {
	go show()
	time.Sleep(time.Second)
}
*/
/* package main

import (
	"fmt"
	"runtime"
)

func show(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("msg: %v\n", msg)
	}
}
func main() {
	go show("java") // 子协程来执行
	for i := 0; i < 2; i++ {
		runtime.Gosched() // 我有权利执行任务，让给（其他子协程来执行）
		fmt.Println("Golang")
	}
	fmt.Println("end...")

}
*/
