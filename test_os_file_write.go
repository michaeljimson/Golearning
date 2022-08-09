/* package main

import "os"

func write() {
	// read write append trunc
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755)
	f.WriteString(" hello mother fucker222")
	f.Close()
}

func writeString() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755)
	f.WriteString("hello python")
	f.Close()
}

func writeAt() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0755)
	f.WriteAt([]byte("aaa"), 3)
	f.Close()
}
func main() {
	//write()
	//writeString()
	writeAt()
}
*/