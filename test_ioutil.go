/* package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func testReadAll() {
	// r := strings.NewReader("hello world!")
	f, _ := os.Open("a.txt")
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(b): %v\n", string(b))
}

func testReadDir() {
	fi, _ := ioutil.ReadDir(".")
	for _, v := range fi {
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

func testReadFile() {
	b, _ := ioutil.ReadFile("a.txt")
	fmt.Printf("string(b): %v\n", string(b))
}

func testWriteFile() {
	ioutil.WriteFile("a.txt", []byte("hello world"), 0664)
}

func testTempFile() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tmpfile.Name(): %v\n", tmpfile.Name())

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
func main() {
	// testReadDir()
	// testReadFile()
	// testWriteFile()
	testTempFile()
}
*/