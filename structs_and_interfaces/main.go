package main

import "fmt"

type Person struct {
	age  uint8
	name string
}

type Child struct {
	age    uint8
	mother Person
	father Person
}

func (c Child) printParents() {
	fmt.Printf("The parents are %v and %v\n", c.mother.name, c.father.name)
}

func (c Child) printAge() {
	fmt.Printf("I am %v years old\n", c.age)
}

func (p Person) printAge() {
	fmt.Printf("I am %v years old\n", p.age)
}

type Ages interface {
	printAge()
}

func ageTester(a Ages) {
	a.printAge()
}

func main() {
	var Peter Person
	fmt.Printf("Peters name is %v and his age is %v\n", Peter.name, Peter.age)

	Peter.age = 40
	Peter.name = "Peter"
	fmt.Printf("Peters name is %v and his age is %v\n", Peter.name, Peter.age)

	var Susanne Person = Person{age: 23, name: "Susanne"}
	fmt.Printf("Susannes name is %v and his age is %v\n", Susanne.name, Susanne.age)

	var billy Child = Child{age: 4, mother: Susanne, father: Person{name: "Bob"}}
	fmt.Printf("Billys parents are %v and %v\n", billy.mother.name, billy.father.name)

	var jeff = struct {
		age  uint8
		name string
	}{80, "Jeff"}
	fmt.Printf("%v is %v years old\n", jeff.name, jeff.age)

	billy.printParents()

	ageTester(billy)
	ageTester(Susanne)
}
