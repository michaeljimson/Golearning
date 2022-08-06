/* package main

import "fmt"

type Person struct {
	id   int
	name string
}

func showPerson(per Person) {
	per.id = 101
	per.name = "jimson"
	fmt.Printf("per: %v\n", per)
}

func showPerson2(per *Person) {
	per.id = 102
	per.name = "Genius"
	fmt.Printf("per: %v\n", per)
}
func main() {
	tom := Person{
		id:   1000,
		name: "michael",
	}

	/* fmt.Printf("tom: %v\n", tom)
	fmt.Println("---------------")
	showPerson(tom)
	fmt.Printf("tom: %v\n", tom) */
	per := &tom
	fmt.Printf("per: %v\n", per)
	showPerson2(per)
	fmt.Println("--------------")
	fmt.Printf("per: %v\n", per)

}
 */