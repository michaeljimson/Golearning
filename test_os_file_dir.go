/* package main

import (
	"fmt"
	"os"
)

// create files

func creasteFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// create dir
func makeDir() {
	/* err := os.Mkdir("a", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} */

	err := os.MkdirAll("a/b/c", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// rm dir file
func remove() {
	/* err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} */
	err := os.RemoveAll("a")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// modify
func wd() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
	os.Chdir("d:/")
	dir, _ = os.Getwd()
	fmt.Printf("dir: %v\n", dir)

	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

// rename
func rename() {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// read file
func readFile() {
	b, _ := os.ReadFile("test2.txt")
	fmt.Printf("b: %v\n", string(b[:]))
}

// write file
func writeFile() {
	os.WriteFile("test2.txt", []byte("hello"), os.ModePerm)
}
func main() {
	//creasteFile()
	//makeDir()
	//remove()
	//wd()
	//rename()
	//readFile()
	writeFile()
}
 */