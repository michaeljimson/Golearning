/* package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	/* counter := 1
	for _ = range ticker.C {
		fmt.Println("ticker...")
		counter++
		if counter > 5 {
			ticker.Stop()
			break
		}
	} */

	chanInt := make(chan int)

	go func() {
		for _ = range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()
	sum := 0
	for v := range chanInt {
		fmt.Printf("收到: %v\n", v)

		if sum >= 10 {
			fmt.Printf("sum: %v\n", sum)
			break
		}
		sum += v

	}
}
 */