package main

import "fmt"

type hasIdentity interface {
	getName() string
}

type Human struct {
	Name string
}

type Animal struct {
	Name string
}

func printHello(name string, human Human) {
	fmt.Println("Hello ", name , ", my name is" , human.getName())
}

func printHelloAnimal(name string, animal Animal){
	fmt.Println("Miaw ", name , ", miaw miaw " , animal.getName())
}

func (human Human) getName() string {
	return human.Name
}

func (animal Animal) getName()string {
	return animal.Name;
}

func main() {
	rizal := Human{Name: "M Rizal F"}
	kucing := Animal{Name: "Boba"}

	printHello("Miftah", rizal)

	printHelloAnimal("Rizal", kucing)
}