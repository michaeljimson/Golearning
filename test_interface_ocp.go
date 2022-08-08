/* package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
}

type Cat struct {
}

func (dog Dog) eat() {
	fmt.Println("Dog eat...")
}
func (dog Dog) sleep() {
	fmt.Println("Dog sleep...")
}

func (cat Cat) eat() {
	fmt.Println("Cat eat...")
}
func (cat Cat) sleep() {
	fmt.Println("Cat sleep...")
}

type Person struct {
}

//pet 既可以传递Dog 也可以传递Cat
func (person Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}
func main() {
	dog := Dog{}
	cat := Cat{}

	person := Person{}
	person.care(dog)
	person.care(cat)

}
*/