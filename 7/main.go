package main

import "fmt"

type Person struct {
	Name string
	Age  uint8
}

type FisicPerson struct {
	Person
	cpf string
}

func (p *FisicPerson) ChangeName(name string) {
	p.Person.Name = name
}

func main() {
	person := FisicPerson{
		Person: Person{
			Name: "Guilherme",
			Age:  26,
		},
		cpf: "439.882.918-03",
	}

	fmt.Println(person)

	person.ChangeName("Marta")

	fmt.Println(person)
}
