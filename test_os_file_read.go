/* package main

import (
	"fmt"
	"os"
)

func openClose() {
	/* // read only
	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	} */

	f, err := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 755)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	}
}

// create File
func createFile() {
	f, _ := os.Create("a2.txt")
	fmt.Printf("f.Name(): %v\n", f.Name())
	f2, _ := os.CreateTemp("", "temp")
	fmt.Printf("f2.Name(): %v\n", f2.Name())
}

// read

func readOps() {
	/* f, _ := os.Open("a.txt")
	for {
		buf := make([]byte, 3)
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(buf): %v\n", string(buf))
	}
	f.Close() */
	/* f, _ := os.Open("a.txt")
	buf := make([]byte, 4)
	n, _ := f.ReadAt(buf, 3)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf)) */

	/* de, _ := os.ReadDir("a")
	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	} */

	f, _ := os.Open("a.txt")
	f.Seek(3, 0)
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f.Close()
}
func main() {
	//openClose()
	//createFile()
	readOps()
}
 */